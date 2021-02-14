package application_test

import (
	"coffee-machine-dz/pkg/coffee-machine/application"
	infra "coffee-machine-dz/pkg/coffee-machine/infrastructure/coffee-machine"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIngredientService(t *testing.T) {
	repo := infra.NewIngredientMemRepo()
	svc := application.NewIngredientService(repo)

	err := svc.Save("hot_milk")
	assert.NoError(t, err)

	_, err = svc.ByName("abc")
	assert.Error(t, err)

	ingd, err := svc.ByName("hot_milk")
	assert.NoError(t, err)
	assert.EqualValues(t, "hot_milk", ingd.Name)
}
