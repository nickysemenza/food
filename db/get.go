package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"
	log "github.com/sirupsen/logrus"
)

// GetRecipeDetailSections finds the sections.
func (c *Client) GetRecipeDetailSections(ctx context.Context, detailID string) ([]Section, error) {
	ctx, span := c.tracer.Start(ctx, "GetRecipeDetailSections")
	defer span.End()
	var res []Section
	if err := c.selectContext(ctx, c.psql.Select("*").From(sectionsTable).Where(sq.Eq{"recipe_detail": detailID}), &res); err != nil {
		return nil, err
	}
	return res, nil
}

// GetSectionInstructions finds the instructions for a section.
func (c *Client) GetSectionInstructions(ctx context.Context, sectionUUID string) ([]SectionInstruction, error) {
	ctx, span := c.tracer.Start(ctx, "GetSectionInstructions")
	defer span.End()
	var res []SectionInstruction
	if err := c.selectContext(ctx, c.psql.Select("*").From(sInstructionsTable).Where(sq.Eq{"section": sectionUUID}), &res); err != nil {
		return nil, err
	}
	return res, nil

}

// GetSectionIngredients finds the ingredients for a section.
func (c *Client) GetSectionIngredients(ctx context.Context, sectionUUID string) ([]SectionIngredient, error) {
	ctx, span := c.tracer.Start(ctx, "GetSectionIngredients")
	defer span.End()
	var res []SectionIngredient
	if err := c.selectContext(ctx, c.psql.Select("*").From(sIngredientsTable).Where(sq.Eq{"section": sectionUUID}), &res); err != nil {
		return nil, err
	}
	return res, nil
}

// GetIngredientByUUID finds an ingredient.
func (c *Client) GetIngredientByUUID(ctx context.Context, uuid string) (*Ingredient, error) {
	ctx, span := c.tracer.Start(ctx, "GetIngredientByUUID")
	defer span.End()
	cacheKey := fmt.Sprintf("i:%s", uuid)

	cval, hit := c.cache.Get(cacheKey)
	log.WithField("key", cacheKey).WithField("hit", hit).Debug("cache:ingredients")
	if hit {
		ing, ok := cval.(Ingredient)
		if ok {
			return &ing, nil
		}
	}

	query, args, err := c.psql.Select("*").From(ingredientsTable).Where(sq.Eq{"uuid": uuid}).ToSql()
	if err != nil {
		return nil, err
	}
	ingredient := &Ingredient{}
	err = c.db.GetContext(ctx, ingredient, query, args...)

	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}
	c.cache.SetWithTTL(cacheKey, *ingredient, 0, time.Second)
	return ingredient, nil
}

// GetRecipeByUUID gets a recipe by name, shallowly.
func (c *Client) GetRecipeDetailWhere(ctx context.Context, eq sq.Eq) (*RecipeDetail, error) {
	ctx, span := c.tracer.Start(ctx, "GetRecipeDetailWhere")
	defer span.End()
	return c.getRecipeDetail(ctx,
		c.psql.Select("*").
			From(recipeDetailsTable).
			Where(eq).
			OrderBy("version DESC").
			Limit(1))
}

func (c *Client) getRecipeDetail(ctx context.Context, sb sq.SelectBuilder) (*RecipeDetail, error) {
	ctx, span := c.tracer.Start(ctx, "getRecipeDetail")
	defer span.End()
	query, args, err := sb.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build query: %w", err)
	}

	r := &RecipeDetail{}
	err = c.db.GetContext(ctx, r, query, args...)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to select: %w", err)
	}
	return r, nil
}

// GetRecipes returns all recipes, shallowly.
func (c *Client) GetRecipes(ctx context.Context, searchQuery string, opts ...SearchOption) ([]RecipeDetail, uint64, error) {
	ctx, span := c.tracer.Start(ctx, "GetRecipes")
	defer span.End()

	q := c.psql.Select("*").From(recipeDetailsTable)
	cq := c.psql.Select("count(*)").From(recipeDetailsTable)
	q = newSearchQuery(opts...).apply(q)
	cq = newSearchQuery(opts...).apply(cq)
	if searchQuery != "" {
		q = q.Where(sq.ILike{"name": fmt.Sprintf("%%%s%%", searchQuery)})
		cq = cq.Where(sq.ILike{"name": fmt.Sprintf("%%%s%%", searchQuery)})
	}
	cq = cq.RemoveLimit().RemoveOffset()

	r := []RecipeDetail{}
	err := c.selectContext(ctx, q, &r)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, 0, nil
	}
	if err != nil {
		return nil, 0, fmt.Errorf("failed to select: %w", err)
	}
	var count uint64
	if err := c.getContext(ctx, cq, &count); err != nil {
		return nil, 0, err
	}
	return r, count, nil
}

// GetRecipeDetailsWithIngredient gets all recipes with an ingredeitn
// todo: consolidate into getrecipes
func (c *Client) GetRecipeDetailsWithIngredient(ctx context.Context, ingredient string) ([]RecipeDetail, error) {
	ctx, span := c.tracer.Start(ctx, "GetRecipesWithIngredient")
	defer span.End()
	query, args, err := c.psql.Select(getRecipeDetailColumns()...).From(recipeDetailsTable).
		Join("recipe_sections on recipe_sections.recipe_detail = recipe_details.uuid").
		Join("recipe_section_ingredients on recipe_sections.uuid = recipe_section_ingredients.section").
		Where(sq.Eq{"ingredient": ingredient}).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build query: %w", err)
	}

	r := []RecipeDetail{}
	err = c.db.SelectContext(ctx, &r, query, args...)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to select: %w", err)
	}
	return r, nil
}

// GetRecipeDetailByUUIDFull gets a recipe by UUID, with all dependencies.
func (c *Client) GetRecipeDetailByUUIDFull(ctx context.Context, uuid string) (*RecipeDetail, error) {
	ctx, span := c.tracer.Start(ctx, "GetRecipeDetailByUUIDFull")
	defer span.End()
	r, err := c.GetRecipeDetailWhere(ctx, sq.Eq{"uuid": uuid})
	if err != nil {
		return nil, err
	}
	if r == nil {
		return r, nil
	}

	r.Sections, err = c.GetRecipeDetailSections(ctx, uuid)
	if err != nil {
		return nil, err
	}

	for x, s := range r.Sections {
		r.Sections[x].Instructions, err = c.GetSectionInstructions(ctx, s.UUID)
		if err != nil {
			return nil, err
		}
		r.Sections[x].Ingredients, err = c.GetSectionIngredients(ctx, s.UUID)
		if err != nil {
			return nil, err
		}

		for y, i := range r.Sections[x].Ingredients {
			if i.IngredientUUID.String != "" {
				ing, err := c.GetIngredientByUUID(ctx, i.IngredientUUID.String)
				if err != nil {
					return nil, err
				}
				r.Sections[x].Ingredients[y].RawIngredient = ing
			}
			if i.RecipeUUID.String != "" {
				rec, err := c.GetRecipeDetailWhere(ctx, sq.Eq{"recipe": i.RecipeUUID.String})
				if err != nil {
					return nil, err
				}
				r.Sections[x].Ingredients[y].RawRecipe = rec
			}
		}
	}

	return r, nil
}
func (c *Client) getIngredients(ctx context.Context, addons func(q sq.SelectBuilder) sq.SelectBuilder) ([]Ingredient, uint64, error) {
	ctx, span := c.tracer.Start(ctx, "getIngredients")
	defer span.End()
	q := addons(c.psql.Select("*").From(ingredientsTable))
	cq := addons(c.psql.Select("count(*)").From(ingredientsTable)).RemoveLimit().RemoveOffset()

	i := []Ingredient{}
	err := c.selectContext(ctx, q, &i)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, 0, nil
	}
	if err != nil {
		return nil, 0, fmt.Errorf("failed to select: %w", err)
	}
	var count uint64
	if err := c.getContext(ctx, cq, &count); err != nil {
		return nil, 0, err
	}

	return i, count, nil
}

type SearchQuery struct {
	offset uint64
	limit  uint64
}

func (s *SearchQuery) apply(q sq.SelectBuilder) sq.SelectBuilder {
	if s.limit != 0 {
		q = q.Limit(s.limit)
	}
	if s.offset != 0 {
		q = q.Offset(s.offset)
	}
	return q
}

type SearchOption func(*SearchQuery)

func WithOffset(offset uint64) SearchOption {
	return func(q *SearchQuery) {
		q.offset = offset
	}
}
func WithLimit(limit uint64) SearchOption {
	return func(q *SearchQuery) {
		q.limit = limit
	}
}

func newSearchQuery(opts ...SearchOption) *SearchQuery {
	q := &SearchQuery{}
	for _, opt := range opts {
		// Call the option giving the instantiated
		// *House as the argument
		opt(q)
	}
	return q
}

// GetIngredients returns all ingredients.
func (c *Client) GetIngredients(ctx context.Context, name string, opts ...SearchOption) ([]Ingredient, uint64, error) {
	ctx, span := c.tracer.Start(ctx, "GetIngredients")
	defer span.End()
	return c.getIngredients(ctx, func(q sq.SelectBuilder) sq.SelectBuilder {
		q = q.Where(sq.Eq{"same_as": nil})
		q = newSearchQuery(opts...).apply(q)
		if name != "" {
			return q.Where(sq.ILike{"name": fmt.Sprintf("%%%s%%", name)})
		}
		return q
	})
}
func (c *Client) GetIngrientsSameAs(ctx context.Context, parent string) ([]Ingredient, uint64, error) {
	ctx, span := c.tracer.Start(ctx, "GetIngrientsSameAs")
	defer span.End()
	return c.getIngredients(ctx, func(q sq.SelectBuilder) sq.SelectBuilder {
		return q.Where(sq.Eq{"same_as": parent})
	})
}

//TODO: non-gql version
// func (c *Client) GetMeals(ctx context.Context, recipe string) ([]*model.Meal, error) {
// 	query, args, err := c.psql.Select("meal_uuid AS uuid", "name", "notion_link AS notionURL").From("meals").
// 		LeftJoin("meal_recipe on meals.uuid = meal_recipe.meal_uuid").
// 		Where(sq.Eq{"recipe_uuid": recipe}).ToSql()
// 	if err != nil {
// 		return nil, err
// 	}
// 	var res []*model.Meal
// 	err = c.db.SelectContext(ctx, &res, query, args...)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return res, nil
// }
