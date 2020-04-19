package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
	"gopkg.in/guregu/null.v3/zero"

	sq "github.com/Masterminds/squirrel"
)

const (
	sIngredientsTable  = "recipe_section_ingredients"
	sInstructionsTable = "recipe_section_instructions"
	sectionsTable      = "recipe_sections"
	recipesTable       = "recipes"
	ingredientsTable   = "ingredients"
)

// Client is a database client
type Client struct {
	db *sqlx.DB
}

// Recipe represents a recipe
type Recipe struct {
	UUID         string      `db:"uuid"`
	Name         string      `db:"name"`
	TotalMinutes zero.Int    `db:"total_minutes"`
	Equipment    zero.String `db:"equipment"`
	Source       zero.String `db:"source"`
	Servings     zero.Int    `db:"servings"`
	Quantity     zero.Int    `db:"quantity"`
	Unit         zero.String `db:"unit"`
	Sections     []Section
}

// Section represents a Section
type Section struct {
	UUID         string   `db:"uuid"`
	RecipeUUID   string   `db:"recipe"`
	Minutes      zero.Int `db:"minutes"`
	Sort         zero.Int `db:"sort"`
	Ingredients  []SectionIngredient
	Instructions []SectionInstruction
}

// SectionIngredient is a foo
type SectionIngredient struct {
	UUID        string      `db:"uuid"`
	SectionUUID string      `db:"section"`
	Sort        zero.Int    `db:"sort"`
	Name        string      // todo: use this to load an Ingredient
	Grams       zero.Float  `db:"grams"`
	Amount      zero.Float  `db:"amount"`
	Unit        zero.String `db:"unit"`
	Adjective   zero.String `db:"adjective"`
	Optional    zero.Bool   `db:"optional"`

	// one of the following:
	RecipeUUID     zero.String `db:"recipe"`
	IngredientUUID zero.String `db:"ingredient"`
}

// SectionInstruction represents a SectionInstruction
type SectionInstruction struct {
	UUID        string   `db:"uuid"`
	Sort        zero.Int `db:"sort"`
	Instruction string   `db:"instruction"`
	SectionUUID string   `db:"section"`
}

// Ingredient is a globally-scoped ingredient
type Ingredient struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
}

var psql = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

// New creates a new Client
func New(db *sqlx.DB) *Client {
	return &Client{db: db}
}

// ConnnectionString returns a DSN
func ConnnectionString(host, user, password, dbname string, port int64) string {
	return fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
}
func setUUID(val string) string {
	if val != "" {
		return val
	}
	u, _ := uuid.NewV4()
	return u.String()
}

// AssignUUIDs adds uuids where missing
func (c *Client) AssignUUIDs(ctx context.Context, r *Recipe) error {
	r.UUID = setUUID(r.UUID)
	for x := range r.Sections {
		r.Sections[x].UUID = setUUID(r.Sections[x].UUID)
		r.Sections[x].RecipeUUID = r.UUID
		for y := range r.Sections[x].Ingredients {
			r.Sections[x].Ingredients[y].UUID = setUUID(r.Sections[x].Ingredients[y].UUID)
			r.Sections[x].Ingredients[y].SectionUUID = r.Sections[x].UUID
			ing, err := c.IngredientByName(ctx, r.Sections[x].Ingredients[y].Name)
			if err != nil {
				return err
			}
			r.Sections[x].Ingredients[y].IngredientUUID = zero.StringFrom(ing.UUID)
		}
		for y := range r.Sections[x].Instructions {
			r.Sections[x].Instructions[y].UUID = setUUID(r.Sections[x].Instructions[y].UUID)
			r.Sections[x].Instructions[y].SectionUUID = r.Sections[x].UUID
		}
	}
	return nil
}

// IngredientByName retrieves an ingredient by name, creating it if it does not exist
func (c *Client) IngredientByName(ctx context.Context, name string) (*Ingredient, error) {
	ingredient := &Ingredient{}
	err := c.db.GetContext(ctx, ingredient, `SELECT * FROM ingredients
	WHERE name = $1 LIMIT 1`, name)
	if err == sql.ErrNoRows {
		_, err = c.db.ExecContext(ctx, `INSERT INTO ingredients (uuid, name) VALUES ($1, $2)`, setUUID(""), name)
		if err != nil {
			return nil, err
		}
		return c.IngredientByName(ctx, name)
	}
	return ingredient, err

}

func (c *Client) updateRecipe(ctx context.Context, tx *sql.Tx, r *Recipe) error {
	query, args, err := psql.
		Update(recipesTable).Where(sq.Eq{"uuid": r.UUID}).Set("name", r.Name).
		Set("total_minutes", r.TotalMinutes).
		Set("unit", r.Unit).
		ToSql()
	if err != nil {
		return fmt.Errorf("failed to build query: %w", err)
	}
	_, err = tx.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}
	c.AssignUUIDs(ctx, r)
	spew.Dump(r)
	//update recipe_section_instructions
	//update recipe_section_ingredients
	//update recipe_sections

	// psql.Delete(sectionInstructionsTable).Where(sq.Eq{""})

	if _, err = tx.ExecContext(ctx, `DELETE FROM recipe_section_instructions WHERE section IN (SELECT uuid from recipe_sections WHERE recipe = $1)`, r.UUID); err != nil {
		return fmt.Errorf("failed to delete instructions: %w", err)
	}
	if _, err = tx.ExecContext(ctx, `DELETE FROM recipe_section_ingredients WHERE section IN (SELECT uuid from recipe_sections WHERE recipe = $1)`, r.UUID); err != nil {
		return fmt.Errorf("failed to delete ingredients: %w", err)
	}
	if _, err = tx.ExecContext(ctx, `DELETE FROM recipe_sections WHERE recipe = $1`, r.UUID); err != nil {
		return fmt.Errorf("failed to delete sections: %w", err)
	}

	if len(r.Sections) == 0 {
		return nil
	}

	sectionInsert := psql.Insert(sectionsTable).Columns("uuid", "recipe", "minutes")

	for _, s := range r.Sections {
		sectionInsert = sectionInsert.Values(s.UUID, s.RecipeUUID, s.Minutes)
	}
	query, args, err = sectionInsert.ToSql()

	if err != nil {
		return fmt.Errorf("failed to build query: %w", err)
	}
	_, err = tx.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	for _, s := range r.Sections {
		if len(s.Instructions) > 0 {
			instructionsInsert := psql.Insert(sInstructionsTable).Columns("uuid", "section", "instruction")
			for _, i := range s.Instructions {
				instructionsInsert = instructionsInsert.Values(i.UUID, i.SectionUUID, i.Instruction)
			}
			if _, err = instructionsInsert.RunWith(tx).Exec(); err != nil {
				return err
			}
		}

		if len(s.Ingredients) > 0 {
			ingredientsInsert := psql.Insert(sIngredientsTable).Columns("uuid", "section", "ingredient", "recipe",
				"grams", "amount", "unit", "adjective", "optional")
			for _, i := range s.Ingredients {
				ingredientsInsert = ingredientsInsert.Values(i.UUID, i.SectionUUID, i.IngredientUUID, i.RecipeUUID,
					i.Grams, i.Amount, i.Unit, i.Adjective, i.Optional)
			}
			if _, err = ingredientsInsert.RunWith(tx).Exec(); err != nil {
				return err
			}
		}
	}
	return nil

}

// UpdateRecipe updates a recipe
func (c *Client) UpdateRecipe(ctx context.Context, r *Recipe) error {
	tx, err := c.db.Begin()
	if err != nil {
		return err
	}
	err = c.updateRecipe(ctx, tx, r)
	if err != nil {
		return err
	}
	return tx.Commit()

}

// InsertRecipe inserts a recipe
func (c *Client) InsertRecipe(ctx context.Context, r *Recipe) (string, error) {
	r.UUID = setUUID(r.UUID)
	query, args, err := psql.
		Insert(recipesTable).Columns("uuid", "name").Values(r.UUID, r.Name).
		ToSql()
	if err != nil {
		return "", fmt.Errorf("failed to build query: %w", err)
	}
	tx, err := c.db.Begin()
	if err != nil {
		return "", err
	}
	_, err = tx.ExecContext(ctx, query, args...)
	if err != nil {
		return "", err
	}
	err = c.updateRecipe(ctx, tx, r)
	if err != nil {
		return "", err
	}
	return r.UUID, tx.Commit()

}
