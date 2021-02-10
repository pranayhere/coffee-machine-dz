package application_test

import (
	"coffee-machine-dz/pkg/coffee-machine/application"
	infra "coffee-machine-dz/pkg/coffee-machine/infrastructure/coffee-machine"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContainerService(t *testing.T) {
	containerSvc := createContainerSvc()

	err := containerSvc.Save(500, 500, "hot_water")
	assert.NoError(t, err)

	_, err = containerSvc.ByName("abc")
	assert.Error(t, err)

	repoContainer, _ := containerSvc.ByName("hot_water")
	assert.EqualValues(t, "hot_water", repoContainer.Ingredient.Name)
}

func TestContainerService_Update(t *testing.T) {
	containerSvc := createContainerSvc()

	err := containerSvc.Save(500, 500, "hot_water")
	assert.NoError(t, err)

	repoContainer, _ := containerSvc.ByName("hot_water")
	repoContainer.Qty = 100
	containerSvc.Update(repoContainer)

	updatedContainer, _ := containerSvc.ByName("hot_water")
	assert.EqualValues(t, 100, updatedContainer.Qty)
}

func createContainerSvc() *application.ContainerService {
	ingdRepo := infra.NewIngredientMemRepo()
	ingdSvc := application.NewIngredientService(ingdRepo)

	containerRepo := infra.NewContainerMemRepo()
	containerSvc := application.NewContainerService(*ingdSvc, containerRepo)

	_ = ingdSvc.Save("hot_water")

	return containerSvc
}
