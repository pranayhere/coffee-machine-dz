package coffee_machine_test

import (
	domain "coffee-machine-dz/pkg/coffee-machine/domain/coffee-machine"
	infra "coffee-machine-dz/pkg/coffee-machine/infrastructure/coffee-machine"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRecipeMemRepo(t *testing.T) {
	repo := infra.NewRecipeMemRepo()

	recipe := addRecipe(t, repo, "hot_coffee", map[string]int{
		"hot_water":    100,
		"ginger_syrup": 30,
	})

	repoRecipe, err := repo.ByName("hot_coffee")
	assert.NoError(t, err)
	assert.EqualValues(t, recipe, repoRecipe)

	_, err = repo.ByName("abc")
	assert.Error(t, err)

}

func addRecipe(t *testing.T, repo *infra.RecipeMemRepo, recipeName string, recipe map[string]int) interface{} {
	contents := make([]domain.Content, 0)
	for k, v := range recipe {
		ingd, _ := domain.NewIngredient(k)
		content, _ := domain.NewContent(*ingd, v)
		contents = append(contents, *content)
	}

	rec, _ := domain.NewRecipe(recipeName, contents)

	err := repo.Save(rec)
	assert.NoError(t, err)

	return rec
}
