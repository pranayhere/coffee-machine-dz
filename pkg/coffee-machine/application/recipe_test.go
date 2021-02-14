package application_test

import (
	"coffee-machine-dz/pkg/coffee-machine/application"
	infra "coffee-machine-dz/pkg/coffee-machine/infrastructure/coffee-machine"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRecipeService(t *testing.T) {
	recipeSvc := createRecipeSvc()

	err := recipeSvc.Save("hot_coffee", map[string]int{
		"hot_water": 100,
	})
	assert.NoError(t, err)

	repoRecipe, _ := recipeSvc.ByName("hot_coffee")
	assert.EqualValues(t, "hot_coffee", repoRecipe.Name)
}

func createRecipeSvc() *application.RecipeService {
	ingdRepo := infra.NewIngredientMemRepo()
	ingdSvc := application.NewIngredientService(ingdRepo)

	recipeRepo := infra.NewRecipeMemRepo()
	recipeSvc := application.NewRecipeService(*ingdSvc, recipeRepo)
	_ = ingdSvc.Save("hot_water")

	return recipeSvc
}
