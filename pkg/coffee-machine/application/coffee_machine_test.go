package application_test

import (
	fixture "coffee-machine-dz/pkg"
	alerting "coffee-machine-dz/pkg/alerting/application"
	app "coffee-machine-dz/pkg/coffee-machine/application"
	cm "coffee-machine-dz/pkg/coffee-machine/infrastructure/coffee-machine"
	"testing"
)

func TestCoffeeMachineService(t *testing.T) {
	coffeeMachine := createCoffeeMachine()

	coffeeMachine.Init(1)

	close(coffeeMachine.Orders)
}

func createCoffeeMachine() *app.CoffeeMachineService {
	ingdRepo := cm.NewIngredientMemRepo()
	ingdSvc := app.NewIngredientService(ingdRepo)

	containerRepo := cm.NewContainerMemRepo()
	containerSvc := app.NewContainerService(ingdSvc, containerRepo)

	if err := fixture.LoadContainer(*containerSvc); err != nil {
		panic(err)
	}

	recipeRepo := cm.NewRecipeMemRepo()
	recipeSvc := app.NewRecipeService(ingdSvc, recipeRepo)

	if err := fixture.LoadRecipe(*recipeSvc); err != nil {
		panic(err)
	}

	alertingSvc := alerting.NewAlertingService()

	return app.NewCoffeeMachineService(ingdSvc, containerSvc, recipeSvc, alertingSvc)
}
