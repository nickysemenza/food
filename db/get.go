package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/nickysemenza/gourd/graph/model"
	log "github.com/sirupsen/logrus"
)

// GetRecipeSections finds the sections.
func (c *Client) GetRecipeSections(ctx context.Context, recipeUUID string) ([]Section, error) {
	var res []Section
	if err := c.selectContext(ctx, c.psql.Select("*").From(sectionsTable).Where(sq.Eq{"recipe": recipeUUID}), &res); err != nil {
		return nil, err
	}
	return res, nil
}

// GetSectionInstructions finds the instructions for a section.
func (c *Client) GetSectionInstructions(ctx context.Context, sectionUUID string) ([]SectionInstruction, error) {
	var res []SectionInstruction
	if err := c.selectContext(ctx, c.psql.Select("*").From(sInstructionsTable).Where(sq.Eq{"section": sectionUUID}), &res); err != nil {
		return nil, err
	}
	return res, nil

}

// GetSectionIngredients finds the ingredients for a section.
func (c *Client) GetSectionIngredients(ctx context.Context, sectionUUID string) ([]SectionIngredient, error) {
	var res []SectionIngredient
	if err := c.selectContext(ctx, c.psql.Select("*").From(sIngredientsTable).Where(sq.Eq{"section": sectionUUID}), &res); err != nil {
		return nil, err
	}
	return res, nil
}

// GetIngredientByUUID finds an ingredient.
func (c *Client) GetIngredientByUUID(ctx context.Context, uuid string) (*Ingredient, error) {
	cacheKey := fmt.Sprintf("i:%s", uuid)

	ctx, span := c.tracer.Start(ctx, "GetIngredientByUUID")
	defer span.End()

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

// GetRecipeByUUID gets a recipe by UUID, shallowly.
func (c *Client) GetRecipeByUUID(ctx context.Context, uuid string) (*Recipe, error) {
	return c.getRecipe(ctx, c.psql.Select("*").From(recipesTable).Where(sq.Eq{"uuid": uuid}))
}

// GetRecipeByUUID gets a recipe by name, shallowly.
func (c *Client) GetRecipeByName(ctx context.Context, name string) (*Recipe, error) {
	return c.getRecipe(ctx, c.psql.Select("*").From(recipesTable).Where(sq.Eq{"name": name}))
}

func (c *Client) getRecipe(ctx context.Context, sb sq.SelectBuilder) (*Recipe, error) {
	query, args, err := sb.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build query: %w", err)
	}

	r := &Recipe{}
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
func (c *Client) GetRecipes(ctx context.Context, searchQuery string, opts ...SearchOption) ([]Recipe, uint64, error) {
	ctx, span := c.tracer.Start(ctx, "GetRecipes")
	defer span.End()

	q := c.psql.Select("*").From(recipesTable)
	cq := c.psql.Select("count(*)").From(recipesTable)
	q = newSearchQuery(opts...).apply(q)
	cq = newSearchQuery(opts...).apply(cq)
	if searchQuery != "" {
		q = q.Where(sq.ILike{"name": fmt.Sprintf("%%%s%%", searchQuery)})
		cq = cq.Where(sq.ILike{"name": fmt.Sprintf("%%%s%%", searchQuery)})
	}
	cq = cq.RemoveLimit().RemoveOffset()

	r := []Recipe{}
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

// GetRecipesWithIngredient gets all recipes with an ingredeitn
// todo: consolidate into getrecipes
func (c *Client) GetRecipesWithIngredient(ctx context.Context, ingredient string) ([]Recipe, error) {
	ctx, span := c.tracer.Start(ctx, "GetRecipesWithIngredient")
	defer span.End()
	query, args, err := c.psql.Select(getRecipeColumns()...).From(recipesTable).
		Join("recipe_sections on recipe_sections.recipe = recipes.uuid").
		Join("recipe_section_ingredients on recipe_sections.uuid = recipe_section_ingredients.section").
		Where(sq.Eq{"ingredient": ingredient}).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build query: %w", err)
	}

	r := []Recipe{}
	err = c.db.SelectContext(ctx, &r, query, args...)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to select: %w", err)
	}
	return r, nil
}

// GetRecipeByUUIDFull gets a recipe by UUID, with all dependencies.
func (c *Client) GetRecipeByUUIDFull(ctx context.Context, uuid string) (*Recipe, error) {
	r, err := c.GetRecipeByUUID(ctx, uuid)
	if err != nil {
		return nil, err
	}
	if r == nil {
		return r, nil
	}

	r.Sections, err = c.GetRecipeSections(ctx, uuid)
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
				rec, err := c.GetRecipeByUUID(ctx, i.RecipeUUID.String)
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

func (c *Client) GetMeals(ctx context.Context, recipe string) ([]*model.Meal, error) {
	query, args, err := c.psql.Select("meal_uuid AS uuid", "name", "notion_link AS notionURL").From("meals").
		LeftJoin("meal_recipe on meals.uuid = meal_recipe.meal_uuid").
		Where(sq.Eq{"recipe_uuid": recipe}).ToSql()
	if err != nil {
		return nil, err
	}
	var res []*model.Meal
	err = c.db.SelectContext(ctx, &res, query, args...)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Client) GetRecipeSource(ctx context.Context, recipeUUID string) (*model.Source, error) {
	query, args, err := c.psql.Select(
		"name", "meta",
	).From("recipe_sources").Where(sq.Eq{"recipe": recipeUUID}).ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build query: %w", err)
	}

	x := &model.Source{}
	err = c.db.GetContext(ctx, x, query, args...)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to select: %w", err)
	}
	return x, nil
}
