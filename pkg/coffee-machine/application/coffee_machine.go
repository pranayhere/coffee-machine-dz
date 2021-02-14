package application

import (
	alerting "coffee-machine-dz/pkg/alerting/application"
	coffee_machine "coffee-machine-dz/pkg/coffee-machine/domain/coffee-machine"
	"errors"
	"fmt"
	"sync"
)

type CoffeeMachineService struct {
	ingdSvc      IngredientService
	containerSvc ContainerService
	recipeSvc    RecipeService
	alertingSvc  alerting.AlertingService

	Orders  chan string      // jobs
	Recipes chan RecipeError // results

	workerWg   sync.WaitGroup
	producerWg sync.WaitGroup
	resultWg   sync.WaitGroup
}

type RecipeError struct {
	recipe *coffee_machine.Recipe
	err    error
}

func NewCoffeeMachineService(ingdSvc IngredientService, containerSvc ContainerService, recipeSvc RecipeService, alertingSvc alerting.AlertingService) *CoffeeMachineService {
	var workerWg sync.WaitGroup
	var producerWg sync.WaitGroup
	var resultWg sync.WaitGroup

	return &CoffeeMachineService{
		ingdSvc:      ingdSvc,
		containerSvc: containerSvc,
		recipeSvc:    recipeSvc,
		alertingSvc:  alertingSvc,

		Orders:  make(chan string, 5),
		Recipes: make(chan RecipeError, 5),

		workerWg:   workerWg,
		producerWg: producerWg,
		resultWg:   resultWg,
	}
}

type CoffeeMachineSvc interface {
	Start()
	MakeDrink(order []string)
	Stop()
}

func (cm *CoffeeMachineService) Start() {
	cm.resultWg.Add(1)
	go cm.result()
	cm.createWorkerPool(3)
}

func (cm *CoffeeMachineService) MakeDrink(order []string) {
	cm.producerWg.Add(1)
	go cm.process(order)
}

func (cm *CoffeeMachineService) process(order []string) {
	for _, s := range order {
		cm.Orders <- s
	}

	cm.producerWg.Done()
}

func (cm *CoffeeMachineService) worker() {
	for order := range cm.Orders {
		recipe, err := cm.recipeSvc.ByName(order)
		if err != nil {
			cm.Recipes <- RecipeError{recipe: nil, err: err}
		} else {
			err = cm.DispenseIngredient(*recipe)
			if err != nil {
				cm.Recipes <- RecipeError{recipe: nil, err: err}
			} else {
				cm.Recipes <- RecipeError{recipe: recipe, err: nil}
			}
		}
	}

	cm.workerWg.Done()
}

func (cm *CoffeeMachineService) result() {
	defer cm.resultWg.Done()

	for recipe := range cm.Recipes {
		fmt.Println("Result Error is : ", recipe.recipe, " err : ", recipe.err)
		if recipe.err != nil {
			cm.alertingSvc.Alert(recipe.err)
		} else {
			fmt.Println("serving you delicious " + recipe.recipe.Name)
		}
	}
}

func (cm *CoffeeMachineService) createWorkerPool(noOfWorkers int) {
	for i := 0; i < noOfWorkers; i++ {
		cm.workerWg.Add(1)
		go cm.worker()
	}
}

func (cm *CoffeeMachineService) Stop() {
	cm.producerWg.Wait()
	close(cm.Orders)

	cm.workerWg.Wait()
	close(cm.Recipes)

	cm.resultWg.Wait()
}

func (cm *CoffeeMachineService) DispenseIngredient(recipe coffee_machine.Recipe) error {
	for _, content := range recipe.Contents {
		container, _ := cm.containerSvc.ByName(content.Ingredient.Name)
		if container.Qty < content.Qty {
			return errors.New("Not enough Ingredient : " + container.Ingredient.Name)
		}
	}

	for _, content := range recipe.Contents {
		container, _ := cm.containerSvc.ByName(content.Ingredient.Name)

		_, err := container.Dispense(content.Qty)
		if err != nil {
			return errors.New("Not enough Ingredient : " + container.Ingredient.Name)
		}

		err = cm.containerSvc.Update(container)
		if err != nil {
			return err
		}
	}

	return nil
}
