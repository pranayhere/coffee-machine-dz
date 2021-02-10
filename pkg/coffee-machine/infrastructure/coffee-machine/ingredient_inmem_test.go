package coffee_machine_test

import (
	domain "coffee-machine-dz/pkg/coffee-machine/domain/coffee-machine"
	infra "coffee-machine-dz/pkg/coffee-machine/infrastructure/coffee-machine"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIngredientMemRepo(t *testing.T) {
	repo := infra.NewIngredientMemRepo()

	ingd := addIngredient(t, repo, "hot_milk")

	// test idempotench
	_ = addIngredient(t, repo, "hot_milk")

	repoIngd, err := repo.ByName("hot_milk")
	assert.NoError(t, err)

	assert.EqualValues(t, ingd, repoIngd)

	_, err = repo.ByName("abc")
	assert.Error(t, err)
}

func addIngredient(t *testing.T, repo *infra.IngredientMemRepo, ingdName string) *domain.Ingredient {
	ingd, err := domain.NewIngredient(ingdName)
	assert.NoError(t, err)

	err = repo.Save(ingd)
	assert.NoError(t, err)

	return ingd
}