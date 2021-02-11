package main

import (
	"coffee-machine-dz/pkg"
	alerting "coffee-machine-dz/pkg/alerting/application"
	app "coffee-machine-dz/pkg/coffee-machine/application"
	cm "coffee-machine-dz/pkg/coffee-machine/infrastructure/coffee-machine"
	"coffee-machine-dz/pkg/common/cmd"
	"log"
)

func main() {
	log.Println("starting coffee-machine")
	ctx := cmd.Context()

	machine := createCoffeeMachine()
	drinks := []string{"hot_coffee", "hot_tea", "black_tea"}

	go func() {
		machine.Start()
		machine.MakeDrink(drinks)
		machine.MakeDrink(drinks)
		machine.MakeDrink(drinks)
		machine.MakeDrink(drinks)
	}()

	<-ctx.Done()

	log.Println("Stopping coffee-machine")
	machine.Stop()
	log.Println("coffee-machine Stopped")
}

func createCoffeeMachine() *app.CoffeeMachineService {
	ingdRepo := cm.NewIngredientMemRepo()
	ingdSvc := app.NewIngredientService(ingdRepo)

	if err := fixture.LoadIngredient(*ingdSvc); err != nil {
		panic(err)
	}

	containerRepo := cm.NewContainerMemRepo()
	containerSvc := app.NewContainerService(*ingdSvc, containerRepo)

	if err := fixture.LoadContainer(*containerSvc); err != nil {
		panic(err)
	}

	recipeRepo := cm.NewRecipeMemRepo()
	recipeSvc := app.NewRecipeService(*ingdSvc, recipeRepo)

	if err := fixture.LoadRecipe(*recipeSvc); err != nil {
		panic(err)
	}

	alertingSvc := alerting.NewAlertingService()

	return app.NewCoffeeMachineService(*ingdSvc, *containerSvc, *recipeSvc, *alertingSvc)
}