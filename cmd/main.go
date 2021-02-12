package main

import (
	"coffee-machine-dz/pkg"
	alerting "coffee-machine-dz/pkg/alerting/application"
	app "coffee-machine-dz/pkg/coffee-machine/application"
	cm "coffee-machine-dz/pkg/coffee-machine/infrastructure/coffee-machine"
	"log"
	"sync"
)

func main() {
	log.Println("starting coffee-machine")
	machine := createCoffeeMachine()
	drinks := []string{"hot_coffee", "hot_tea", "black_tea"}

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		machine.Start()
		machine.MakeDrink(drinks)
		machine.Stop()

		wg.Done()
	}()

	wg.Wait()
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