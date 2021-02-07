package main

import (
	"coffee-machine-dz/pkg"
	app "coffee-machine-dz/pkg/coffee-machine/application"
	cm "coffee-machine-dz/pkg/coffee-machine/infrastructure/coffee-machine"
	"log"
)

func main() {
	log.Println("starting coffee-machine")

	//ctx := cmd.Context()

	createCoffeeMachine()
	//cm := createCoffeeMachine()

	//drinks := []string{"coffee", "tea", "coffee"}
	//cm.Start()
	//
	//go func() {
	//	cm.MakeDrink(drinks)
	//	cm.MakeDrink(drinks)
	//	cm.MakeDrink(drinks)
	//
	//	cm.Stop()
	//}()
	//
	//<-ctx.Done()
	//log.Println("Stopping coffee-machine")
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

	//ingd, err := ingdSvc.ByName("hot_water")
	//if err != nil {
	//	panic(err)
	//}
	//
	//fmt.Println(ingd.Name)

	return app.NewCoffeeMachineService()
}