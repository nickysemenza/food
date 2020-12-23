package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/nickysemenza/gourd/common"
)

// IngredientByName retrieves an ingredient by name, creating it if it does not exist.
func (c *Client) IngredientByName(ctx context.Context, name string) (*Ingredient, error) {
	ingredient := &Ingredient{}
	err := c.db.GetContext(ctx, ingredient, `SELECT * FROM ingredients
	WHERE name = $1 LIMIT 1`, name)
	if errors.Is(err, sql.ErrNoRows) {
		_, err = c.db.ExecContext(ctx, `INSERT INTO ingredients (id, name) VALUES ($1, $2)`, common.UUID(), name)
		if err != nil {
			return nil, err
		}
		return c.IngredientByName(ctx, name)
	}
	return ingredient, err
}

//nolint: funlen
func (c *Client) updateRecipe(ctx context.Context, tx *sql.Tx, r *RecipeDetail) error {
	q := c.psql.
		Update(recipeDetailsTable).Where(sq.Eq{"id": r.Id}).Set("name", r.Name).
		Set("total_minutes", r.TotalMinutes).
		Set("unit", r.Unit)

	_, err := c.execTx(ctx, tx, q)
	if err != nil {
		return err
	}

	if err := c.AssignIds(ctx, r); err != nil {
		return err
	}

	if len(r.Sections) == 0 {
		return nil
	}

	// sections
	sectionInsert := c.psql.Insert(sectionsTable).Columns("id", "recipe_detail", "minutes")
	for _, s := range r.Sections {
		sectionInsert = sectionInsert.Values(s.Id, s.RecipeDetailId, s.Minutes)
	}

	_, err = c.execTx(ctx, tx, sectionInsert)
	if err != nil {
		return err
	}

	instructionsInsert := c.psql.Insert(sInstructionsTable).Columns("id", "section", "instruction")
	ingredientsInsert := c.psql.Insert(sIngredientsTable).Columns("id", "section", "ingredient", "recipe",
		"grams", "amount", "unit", "adjective", "optional")

	var hasInstructions, hasIngredients bool
	for _, s := range r.Sections {
		for _, i := range s.Instructions {
			hasInstructions = true
			instructionsInsert = instructionsInsert.Values(i.Id, i.SectionId, i.Instruction)
		}

		for _, i := range s.Ingredients {
			hasIngredients = true

			ingredientsInsert = ingredientsInsert.Values(i.Id, i.SectionId, i.IngredientId, i.RecipeId,
				i.Grams, i.Amount, i.Unit, i.Adjective, i.Optional)
		}

	}
	if hasInstructions {
		if _, err = c.execTx(ctx, tx, instructionsInsert); err != nil {
			return err
		}
	}
	if hasIngredients {
		if _, err = c.execTx(ctx, tx, ingredientsInsert); err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) insertRecipe(ctx context.Context, r *RecipeDetail) (*RecipeDetail, error) {
	ctx, span := c.tracer.Start(ctx, "insertRecipe")
	defer span.End()
	if r.Id == "" {
		// todo? needed?
		r.Id = common.UUID()
	}
	log.Println("inserting", r.Id, r.Name)
	q := c.psql.
		Insert(recipeDetailsTable).
		Columns("id", "recipe", "name", "version").
		Values(r.Id, r.RecipeId, r.Name, r.Version)

	tx, err := c.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	_, err = c.execTx(ctx, tx, q)
	if err != nil {
		return nil, err
	}

	err = c.updateRecipe(ctx, tx, r)
	if err != nil {
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return c.GetRecipeDetailByIdFull(ctx, r.Id)
}

// InsertRecipe inserts a recipe.
func (c *Client) InsertRecipe(ctx context.Context, r *RecipeDetail) (*RecipeDetail, error) {
	ctx, span := c.tracer.Start(ctx, "InsertRecipe")
	defer span.End()

	// if we have an existing recipe with the same Id or name, this one is a n+1 version of that one
	version := int64(1)
	var modifying *RecipeDetail
	parentID := ""
	var err error
	if r.Id != "" {
		modifying, err = c.GetRecipeDetailWhere(ctx, sq.Eq{"id": r.Id})
		if err != nil {
			return nil, fmt.Errorf("failed to find prior recipe: %w", err)
		}
	}
	if modifying == nil {
		modifying, err = c.GetRecipeDetailWhere(ctx, sq.Eq{"name": r.Name})
		if err != nil {
			return nil, fmt.Errorf("failed to find prior recipe: %w", err)
		}
	}

	if modifying != nil {
		latestVersion, err := c.GetRecipeDetailWhere(ctx, sq.Eq{"recipe": modifying.RecipeId})
		if err != nil {
			return nil, fmt.Errorf("failed to find prior recipe: %w", err)
		}
		version = latestVersion.Version + 1
		parentID = latestVersion.RecipeId
	}
	r.Version = version
	r.Id = common.UUID()

	if parentID == "" {
		parentID = common.UUID()
		_, err = c.execContext(ctx, c.psql.
			Insert(recipesTable).Columns("id").Values(parentID))
		if err != nil {
			return nil, fmt.Errorf("failed to insert parent recipe: %w", err)
		}
	}
	r.RecipeId = parentID

	return c.insertRecipe(ctx, r)

}

func (c *Client) IngredientToRecipe(ctx context.Context, ingredientID string) (*RecipeDetail, error) {
	ctx, span := c.tracer.Start(ctx, "IngredientToRecipe")
	defer span.End()

	i, err := c.GetIngredientById(ctx, ingredientID)
	if err != nil {
		return nil, err
	}
	if i == nil {
		return nil, fmt.Errorf("failed to find ingredient with id %s", ingredientID)
	}

	newRecipe, err := c.InsertRecipe(ctx, &RecipeDetail{Name: i.Name})
	if err != nil {
		return nil, err
	}

	if _, err = c.execContext(ctx,
		c.psql.
			Update(sIngredientsTable).
			Set("ingredient", nil).
			Set("recipe", newRecipe.RecipeId).
			Where(sq.Eq{"ingredient": ingredientID})); err != nil {
		return nil, fmt.Errorf("failed to update references to transformed ingredient: %w", err)
	}

	if _, err = c.execContext(ctx, c.psql.
		Update(ingredientsTable).
		Set("name", fmt.Sprintf("[deprecated] %s", i.Name)).
		Where(sq.Eq{"id": i.Id})); err != nil {
		return nil, fmt.Errorf("failed to deprecated ingredient: %w", err)
	}

	return newRecipe, nil
}