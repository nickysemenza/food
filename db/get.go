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

type RecipeDetails []RecipeDetail

func (r RecipeDetails) ByDetailId() map[string]RecipeDetail {
	m := make(map[string]RecipeDetail)
	for _, x := range r {
		m[x.Id] = x
	}
	return m
}
func (r RecipeDetails) ByRecipeId() map[string]RecipeDetail {
	m := make(map[string]RecipeDetail)
	for _, x := range r {
		m[x.RecipeId] = x
	}
	return m
}
func (r RecipeDetails) First() *RecipeDetail {
	if len(r) == 0 {
		return nil
	}
	return &r[0]
}

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
func (c *Client) GetSectionInstructions(ctx context.Context, sectionId []string) (map[string][]SectionInstruction, error) {
	ctx, span := c.tracer.Start(ctx, "GetSectionInstructions")
	defer span.End()
	var res []SectionInstruction
	if err := c.selectContext(ctx, c.psql.Select("*").From(sInstructionsTable).Where(sq.Eq{"section": sectionId}), &res); err != nil {
		return nil, err
	}
	byId := make(map[string][]SectionInstruction)
	for _, i := range res {
		byId[i.SectionId] = append(byId[i.SectionId], i)
	}
	return byId, nil

}

// GetSectionIngredients finds the ingredients for a section.
func (c *Client) GetSectionIngredients(ctx context.Context, sectionId []string) (map[string][]SectionIngredient, error) {
	ctx, span := c.tracer.Start(ctx, "GetSectionIngredients")
	defer span.End()

	var res []SectionIngredient
	if err := c.selectContext(ctx, c.psql.Select("*").From(sIngredientsTable).Where(sq.Eq{"section": sectionId}), &res); err != nil {
		return nil, err
	}
	byId := make(map[string][]SectionIngredient)
	for _, i := range res {
		byId[i.SectionId] = append(byId[i.SectionId], i)
	}
	return byId, nil
}

// GetIngredientById finds an ingredient.
func (c *Client) GetIngredientById(ctx context.Context, id string) (*Ingredient, error) {
	ctx, span := c.tracer.Start(ctx, "GetIngredientById")
	defer span.End()
	cacheKey := fmt.Sprintf("i:%s", id)

	cval, hit := c.cache.Get(cacheKey)
	log.WithField("key", cacheKey).WithField("hit", hit).Debug("cache:ingredients")
	if hit {
		ing, ok := cval.(Ingredient)
		if ok {
			return &ing, nil
		}
	}

	query, args, err := c.psql.Select("*").From(ingredientsTable).Where(sq.Eq{"id": id}).ToSql()
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

func (c *Client) getIngredientsById(ctx context.Context, id ...string) (map[string]Ingredient, error) {
	ctx, span := c.tracer.Start(ctx, "getIngredientsById")
	defer span.End()

	var i []Ingredient
	err := c.selectContext(ctx, c.psql.Select("*").From(ingredientsTable).Where(sq.Eq{"id": id}), &i)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("failed to select: %w", err)
	}
	byId := map[string]Ingredient{}
	for _, ing := range i {
		byId[ing.Id] = ing
	}
	return byId, nil
}

// GetRecipeById gets a recipe by name, shallowly.
func (c *Client) GetRecipeDetailWhere(ctx context.Context, eq sq.Sqlizer) (RecipeDetails, error) {
	ctx, span := c.tracer.Start(ctx, "GetRecipeDetailWhere")
	defer span.End()
	return c.getRecipeDetail(ctx,
		c.psql.Select("*").
			From(recipeDetailsTable).
			Where(eq).
			OrderBy("version DESC"),
	)
}

func (c *Client) getRecipeDetail(ctx context.Context, sb sq.SelectBuilder) ([]RecipeDetail, error) {
	ctx, span := c.tracer.Start(ctx, "getRecipeDetail")
	defer span.End()

	r := []RecipeDetail{}

	if err := c.selectContext(ctx, sb, &r); err != nil && !errors.Is(err, sql.ErrNoRows) {
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
		Join("recipe_sections on recipe_sections.recipe_detail = recipe_details.id").
		Join("recipe_section_ingredients on recipe_sections.id = recipe_section_ingredients.section").
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

// GetRecipeDetailByIdFull gets a recipe by Id, with all dependencies.
func (c *Client) GetRecipeDetailByIdFull(ctx context.Context, detailId string) (*RecipeDetail, error) {
	ctx, span := c.tracer.Start(ctx, "GetRecipeDetailByIdFull")
	defer span.End()
	topLevelDetail := sq.Eq{"id": detailId}

	// rs, err := c.GetRecipeDetailWhere(ctx, sq.Eq{"id": id})
	// if err != nil {
	// 	return nil, err
	// }
	// if len(rs) != 1 {
	// 	return nil, fmt.Errorf("failed to find 1 recipe detail with id %s, found %d", id, len(rs))
	// }
	// r := rs[0]

	sections, err := c.GetRecipeDetailSections(ctx, detailId)
	if err != nil {
		return nil, err
	}

	var sectionIds []string
	for _, s := range sections {
		sectionIds = append(sectionIds, s.Id)
	}
	sIns, err := c.GetSectionInstructions(ctx, sectionIds)
	if err != nil {
		return nil, err
	}

	sIng, err := c.GetSectionIngredients(ctx, sectionIds)
	if err != nil {
		return nil, err
	}

	var ingredientIds []string
	var recipeIds []string
	for x, s := range sections {
		for sectionId, i := range sIns {
			if s.Id == sectionId {
				sections[x].Instructions = i
			}
		}
		for sectionId, i := range sIng {
			if s.Id == sectionId {
				sections[x].Ingredients = i
			}
		}
		for _, i := range sections[x].Ingredients {
			if i.IngredientId.String != "" {
				ingredientIds = append(ingredientIds, i.IngredientId.String)
			}
			if i.RecipeId.String != "" {
				recipeIds = append(recipeIds, i.RecipeId.String)
				// rec, err := c.GetRecipeDetailWhere(ctx, sq.Eq{"recipe": i.RecipeId.String})
				// if err != nil {
				// 	return nil, err
				// }
				// sections[x].Ingredients[y].RawRecipe = &rec[0]
			}
		}
	}
	// load  all ingredients from all sections in one go
	ingredientsById, err := c.getIngredientsById(ctx, ingredientIds...)
	if err != nil {
		return nil, err
	}

	var eq sq.Sqlizer
	if len(recipeIds) == 0 {
		eq = topLevelDetail
	} else {
		eq = sq.Or{topLevelDetail, sq.Eq{"recipe": recipeIds}}
	}
	recipes, err := c.GetRecipeDetailWhere(ctx, eq)
	if err != nil {
		return nil, err
	}
	r, ok := recipes.ByDetailId()[detailId]
	if !ok {
		return nil, fmt.Errorf("failed to find recipe with detail id %s", detailId)
	}

	recipesUsed := recipes.ByRecipeId()
	for x := range sections {
		for y, i := range sections[x].Ingredients {
			if i.IngredientId.String != "" {
				res := ingredientsById[i.IngredientId.String]
				sections[x].Ingredients[y].RawIngredient = &res
			}
			if i.RecipeId.String != "" {
				res := recipesUsed[i.RecipeId.String]
				sections[x].Ingredients[y].RawRecipe = &res
			}
		}
	}
	r.Sections = sections

	return &r, nil
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
// 	query, args, err := c.psql.Select("meal_id AS id", "name", "notion_link AS notionURL").From("meals").
// 		LeftJoin("meal_recipe on meals.id = meal_recipe.meal_id").
// 		Where(sq.Eq{"recipe_id": recipe}).ToSql()
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
