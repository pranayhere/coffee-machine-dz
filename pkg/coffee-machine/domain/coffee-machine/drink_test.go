package coffee_machine_test

import (
	cm "coffee-machine-dz/pkg/coffee-machine/domain/coffee-machine"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewContent(t *testing.T) {
	ingd := CreateIngredient(t)

	content, err := cm.NewContent(*ingd, 100)
	assert.NoError(t, err)

	assert.EqualValues(t, ingd, &content.Ingredient)
}

func TestNewRecipe(t *testing.T) {
	 contents := createRecipeContent(t)

	 recipe, err := cm.NewRecipe("coffee", contents)
	 assert.NoError(t, err)

	 assert.EqualValues(t, "coffee", recipe.Name)
	 assert.EqualValues(t, contents, recipe.Contents)
}

func createRecipeContent(t *testing.T) []cm.Content {
	ingds := []string{"hot_milk", "hot_water"}
	contents := make([]cm.Content, 0)

	for _, ingdName := range ingds {
		ingd, err := cm.NewIngredient(ingdName)
		assert.NoError(t, err)

		content, err := cm.NewContent(*ingd, 500)
		assert.NoError(t, err)

		contents = append(contents, *content)
	}

	return contents
}