package api

import (
	"context"
	"fmt"
	"testing"

	"github.com/nickysemenza/gourd/db"
	"github.com/nickysemenza/gourd/manager"
	"github.com/stretchr/testify/require"
)

func TestRecipeFromFile(t *testing.T) {
	require := require.New(t)
	ctx := context.Background()

	r, err := RecipeFromFile(ctx, "../testdata/cookies_1.json")
	require.NoError(err)
	require.Equal("cookies 1", r.Recipe.Name)
	baseName := fmt.Sprintf("%s-%s", r.Recipe.Name, db.GetUUID())
	r.Recipe.Name = baseName

	tdb := db.NewDB(t)
	m := manager.New(tdb, nil, nil)
	apiManager := NewAPI(m)

	r2, err := apiManager.CreateRecipe(ctx, r)
	require.NoError(err)

	require.Equal(baseName, r2.Recipe.Name)
	r.Recipe.Id = "" // reset so we create a dup instead of update, ptr

	r3, err := apiManager.CreateRecipe(ctx, r)
	require.NoError(err)

	require.Equal(fmt.Sprintf("%s (dup)", baseName), r3.Recipe.Name)
}
