package application_test

import (
	alerting "coffee-machine-dz/pkg/alerting/application"
	app "coffee-machine-dz/pkg/coffee-machine/application"
	cm "coffee-machine-dz/pkg/coffee-machine/infrastructure/coffee-machine"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestContainer struct {
	Cap  int
	Qty  int
	Ingd string
}

type TestRecipe struct {
	Name   string
	Recipe map[string]int
}

func TestCoffeeMachineService(t *testing.T) {
	testCases := []struct {
		TestName string

		Order     string
		Container TestContainer
		Recipe    TestRecipe

		ExpectedErr bool
		Error   error
	}{
		{
			TestName:    "Prepared: Hot Coffee",
			Order:       "hot_coffee",
			Container:   TestContainer{100, 100, "hot_water"},
			Recipe:      TestRecipe{"hot_coffee", map[string]int{"hot_water": 100}},
			ExpectedErr: false,
			Error: nil,
		},
		{
			TestName: "Ingredient Not Available",
			Order:    "hot_tea",
			Container: TestContainer{100, 100, "hot_tea"},
			Recipe:      TestRecipe{"hot_tea", map[string]int{"ginger_syrup": 100}},
			ExpectedErr: true,
			Error: errors.New("hot_tea cannot be prepared because ginger_syrup is not available"),
		},
		{
			TestName: "Insufficient Ingredients",
			Order:    "hot_tea",
			Container: TestContainer{100, 50, "ginger_syrup"},
			Recipe:      TestRecipe{"hot_tea", map[string]int{"ginger_syrup": 100}},
			ExpectedErr: true,
			Error: errors.New("hot_tea cannot be prepared because item ginger_syrup is not sufficient"),
		},
	}
	for _, c := range testCases {
		t.Run(c.TestName, func(t *testing.T) {
			m := createCoffeeMachine().Init(1)

			_ = m.ContainerSvc.Save(c.Container.Cap, c.Container.Qty, c.Container.Ingd)
			_ = m.RecipeSvc.Save(c.Recipe.Name, c.Recipe.Recipe)

			m.Orders <- c.Order
			close(m.Orders)

			for o := range m.Orders {
				recipe, _ := m.RecipeSvc.ByName(o)
				_, err := m.DispenseIngredient(*recipe)

				if !c.ExpectedErr {
					assert.NoError(t, err)
				} else {
					assert.Equal(t, c.Error, err)
				}
			}
		})
	}
}

func createCoffeeMachine() *app.CoffeeMachineService {
	ingdRepo := cm.NewIngredientMemRepo()
	ingdSvc := app.NewIngredientService(ingdRepo)

	containerRepo := cm.NewContainerMemRepo()
	containerSvc := app.NewContainerService(ingdSvc, containerRepo)

	recipeRepo := cm.NewRecipeMemRepo()
	recipeSvc := app.NewRecipeService(ingdSvc, recipeRepo)

	alertingSvc := alerting.NewAlertingService()

	return app.NewCoffeeMachineService(ingdSvc, containerSvc, recipeSvc, alertingSvc)
}