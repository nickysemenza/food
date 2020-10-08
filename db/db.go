package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/dgraph-io/ristretto"
	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
	"go.opentelemetry.io/otel/api/global"
	"go.opentelemetry.io/otel/api/trace"
	"gopkg.in/guregu/null.v3/zero"

	sq "github.com/Masterminds/squirrel"
)

const (
	sIngredientsTable  = "recipe_section_ingredients"
	sInstructionsTable = "recipe_section_instructions"
	sectionsTable      = "recipe_sections"
	recipesTable       = "recipes"
	ingredientsTable   = "ingredients"
	sourcesTable       = "recipe_sources"
)

// Client is a database client
type Client struct {
	db     *sqlx.DB
	psql   sq.StatementBuilderType
	cache  *ristretto.Cache
	tracer trace.Tracer
}

func getRecipeColumns() []string {
	return []string{"recipes.uuid", "name", "total_minutes", "equipment", "source", "servings", "quantity", "recipes.unit"}
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
	Sources      []Source
}
type Source struct {
	Name string `db:"name"`
	Meta string `db:"meta"`
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
	Grams       zero.Float  `db:"grams"`
	Amount      zero.Float  `db:"amount"`
	Unit        zero.String `db:"unit"`
	Adjective   zero.String `db:"adjective"`
	Optional    zero.Bool   `db:"optional"`

	// one of the following is required for get and update:
	RecipeUUID     zero.String `db:"recipe"`
	IngredientUUID zero.String `db:"ingredient"`

	// one of these is populated via gets
	RawRecipe     *Recipe
	RawIngredient *Ingredient

	// deprecated
	// Name        string      // todo: use this to load an Ingredient
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
	UUID   string      `json:"uuid"`
	Name   string      `json:"name"`
	FdcID  zero.Int    `db:"fdc_id"`
	SameAs zero.String `db:"same_as"`
}

// New creates a new Client.
func New(dbConn *sql.DB) (*Client, error) {
	dbx := sqlx.NewDb(dbConn, "postgres")
	if err := dbx.Ping(); err != nil {
		return nil, err
	}

	// nolint:gomnd
	cache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e7,     // number of keys to track frequency of (10M).
		MaxCost:     1 << 30, // maximum cost of cache (1GB).
		BufferItems: 64,      // number of keys per Get buffer.
		Metrics:     true,
	})
	if err != nil {
		return nil, err
	}
	return &Client{
		db:     dbx,
		psql:   sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
		cache:  cache,
		tracer: global.Tracer("db"),
	}, nil
}

// ConnnectionString returns a DSN.
func ConnnectionString(host, user, password, dbname string, port int64) string {
	return fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
}
func setUUID(val string) string {
	if val != "" {
		return val
	}
	return getUUID()
}
func getUUID() string {
	u, _ := uuid.NewV4()
	return u.String()
}

// AssignUUIDs adds uuids where missing.
func (c *Client) AssignUUIDs(ctx context.Context, r *Recipe) error {
	r.UUID = setUUID(r.UUID)
	for x := range r.Sections {
		r.Sections[x].UUID = setUUID(r.Sections[x].UUID)
		r.Sections[x].RecipeUUID = r.UUID
		for y := range r.Sections[x].Ingredients {
			r.Sections[x].Ingredients[y].UUID = setUUID(r.Sections[x].Ingredients[y].UUID)
			r.Sections[x].Ingredients[y].SectionUUID = r.Sections[x].UUID

			// ingName := r.Sections[x].Ingredients[y].Name
			// if strings.HasPrefix(ingName, "r:") {
			// 	recipeName := strings.TrimPrefix(ingName, "r:")
			// 	recipe, err := c.GetRecipeByName(ctx, recipeName)
			// 	if err != nil {
			// 		return err
			// 	}
			// 	if recipe == nil {
			// 		return fmt.Errorf("no recipe with name %s", recipeName)
			// 	}
			// 	r.Sections[x].Ingredients[y].RecipeUUID = zero.StringFrom(recipe.UUID)
			// } else {
			// 	ing, err := c.IngredientByName(ctx, ingName)
			// 	if err != nil {
			// 		return err
			// 	}
			// 	r.Sections[x].Ingredients[y].IngredientUUID = zero.StringFrom(ing.UUID)
			// }
		}
		for y := range r.Sections[x].Instructions {
			r.Sections[x].Instructions[y].UUID = setUUID(r.Sections[x].Instructions[y].UUID)
			r.Sections[x].Instructions[y].SectionUUID = r.Sections[x].UUID
		}
	}
	return nil
}

// IngredientByName retrieves an ingredient by name, creating it if it does not exist.
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

//nolint: funlen
func (c *Client) updateRecipe(ctx context.Context, tx *sql.Tx, r *Recipe) error {
	query, args, err := c.psql.
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
	if err := c.AssignUUIDs(ctx, r); err != nil {
		return err
	}
	//update recipe_section_instructions
	//update recipe_section_ingredients
	//update recipe_sections

	// c.psql.Delete(sectionInstructionsTable).Where(sq.Eq{""})

	if _, err = tx.ExecContext(ctx,
		`DELETE FROM recipe_section_instructions 
		WHERE section 
		IN (SELECT uuid from recipe_sections WHERE recipe = $1)`, r.UUID); err != nil {
		return fmt.Errorf("failed to delete instructions: %w", err)
	}
	if _, err = tx.ExecContext(ctx,
		`DELETE FROM recipe_section_ingredients 
		WHERE section 
		IN (SELECT uuid from recipe_sections WHERE recipe = $1)`, r.UUID); err != nil {
		return fmt.Errorf("failed to delete ingredients: %w", err)
	}
	if _, err = tx.ExecContext(ctx,
		`DELETE FROM recipe_sections WHERE recipe = $1`, r.UUID); err != nil {
		return fmt.Errorf("failed to delete sections: %w", err)
	}

	if _, err = tx.ExecContext(ctx,
		`DELETE FROM recipe_sources WHERE recipe = $1`, r.UUID); err != nil {
		return fmt.Errorf("failed to delete sources: %w", err)
	}

	if len(r.Sections) == 0 {
		return nil
	}

	// sections
	sectionInsert := c.psql.Insert(sectionsTable).Columns("uuid", "recipe", "minutes")
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

	if len(r.Sources) > 0 {
		sourcesInsert := c.psql.Insert(sourcesTable).Columns("uuid", "recipe", "name", "meta")
		for _, s := range r.Sources {
			sourcesInsert = sourcesInsert.Values(getUUID(), r.UUID, s.Name, s.Meta)
		}
		query, args, err = sourcesInsert.ToSql()

		if err != nil {
			return fmt.Errorf("failed to build query: %w", err)
		}
		_, err = tx.ExecContext(ctx, query, args...)
		if err != nil {
			return err
		}
	}

	for _, s := range r.Sections {
		if len(s.Instructions) > 0 {
			instructionsInsert := c.psql.Insert(sInstructionsTable).Columns("uuid", "section", "instruction")
			for _, i := range s.Instructions {
				instructionsInsert = instructionsInsert.Values(i.UUID, i.SectionUUID, i.Instruction)
			}
			if _, err = instructionsInsert.RunWith(tx).ExecContext(ctx); err != nil {
				return err
			}
		}

		if len(s.Ingredients) > 0 {
			ingredientsInsert := c.psql.Insert(sIngredientsTable).Columns("uuid", "section", "ingredient", "recipe",
				"grams", "amount", "unit", "adjective", "optional")
			for _, i := range s.Ingredients {
				ingredientsInsert = ingredientsInsert.Values(i.UUID, i.SectionUUID, i.IngredientUUID, i.RecipeUUID,
					i.Grams, i.Amount, i.Unit, i.Adjective, i.Optional)
			}
			if _, err = ingredientsInsert.RunWith(tx).ExecContext(ctx); err != nil {
				return err
			}
		}
	}
	return nil
}

// UpdateRecipe updates a recipe.
func (c *Client) UpdateRecipe(ctx context.Context, r *Recipe) error {
	tx, err := c.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	err = c.updateRecipe(ctx, tx, r)
	if err != nil {
		return err
	}
	return tx.Commit()
}

// InsertRecipe inserts a recipe.
func (c *Client) InsertRecipe(ctx context.Context, r *Recipe) (string, error) {
	r.UUID = setUUID(r.UUID)
	query, args, err := c.psql.
		Insert(recipesTable).Columns("uuid", "name").Values(r.UUID, r.Name).
		ToSql()
	if err != nil {
		return "", fmt.Errorf("failed to build query: %w", err)
	}
	tx, err := c.db.BeginTx(ctx, nil)
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

func (c *Client) getContext(ctx context.Context, q sq.SelectBuilder, dest interface{}) error {
	ctx, span := c.tracer.Start(ctx, "getContext")
	defer span.End()

	query, args, err := q.ToSql()
	if err != nil {
		return fmt.Errorf("failed to build query: %w", err)
	}
	err = c.db.GetContext(ctx, dest, query, args...)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf("failed to GetContext: %w", err)
	}
	return nil
}
func (c *Client) selectContext(ctx context.Context, q sq.SelectBuilder, dest interface{}) error {
	ctx, span := c.tracer.Start(ctx, "selectContext")
	defer span.End()

	query, args, err := q.ToSql()
	if err != nil {
		return fmt.Errorf("failed to build query: %w", err)
	}
	err = c.db.SelectContext(ctx, dest, query, args...)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf("failed to SelectContext: %w", err)
	}
	return nil
}

func (c *Client) execContext(ctx context.Context, q sq.InsertBuilder) (sql.Result, error) {
	ctx, span := c.tracer.Start(ctx, "execContext")
	defer span.End()

	query, args, err := q.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build query: %w", err)
	}
	return c.db.ExecContext(ctx, query, args...)
}

func (c *Client) GetKV(ctx context.Context, key string) (string, error) {
	var json string
	q := c.psql.Select("value").From("kv").Where(sq.Eq{"key": key})
	err := c.getContext(ctx, q, &json)
	return json, err
}
func (c *Client) SetKV(ctx context.Context, key string, json string) error {
	q := c.psql.Insert("kv").Columns("key", "value").Values(key, json).Suffix("ON CONFLICT (key) DO UPDATE SET value = ?", json)
	_, err := c.execContext(ctx, q)
	return err
}

func (c *Client) GetAlbums(ctx context.Context) ([]string, error) {
	var albums []string
	q := c.psql.Select("id").From("gphotos_albums")
	err := c.selectContext(ctx, q, &albums)
	return albums, err
}

type Photo struct {
	AlbumID string    `db:"album_id"`
	PhotoID string    `db:"id"`
	Created time.Time `db:"creation_time"`
}

func (c *Client) UpsertPhotos(ctx context.Context, photos []Photo) error {
	q := c.psql.Insert("gphotos_photos").Columns("id", "album_id", "creation_time")
	for _, photo := range photos {
		q = q.Values(photo.PhotoID, photo.AlbumID, photo.Created)
	}
	q = q.Suffix("ON CONFLICT (id) DO NOTHING")
	_, err := c.execContext(ctx, q)
	return err
}

func (c *Client) GetPhotos(ctx context.Context) ([]Photo, error) {
	q := c.psql.Select("*").From("gphotos_photos").OrderBy("creation_time DESC")
	var results []Photo
	err := c.selectContext(ctx, q, &results)
	return results, err

}
