package application

import (
	coffee_machine "coffee-machine-dz/pkg/coffee-machine/domain/coffee-machine"
	"errors"
	"fmt"
	"sync"
)

type CoffeeMachineService struct {
	ingdSvc IngredientService
	containerSvc ContainerService
	recipeSvc RecipeService

	jobs chan string
	results chan RecipeError
	workerWg sync.WaitGroup
	producerWg sync.WaitGroup
}

type RecipeError struct {
	recipe *coffee_machine.Recipe
	err error
}

func NewCoffeeMachineService(ingdSvc IngredientService, containerSvc ContainerService, recipeSvc RecipeService) *CoffeeMachineService {
	var workerWg sync.WaitGroup
	var producerWg sync.WaitGroup

	return &CoffeeMachineService{
		ingdSvc: ingdSvc,
		containerSvc: containerSvc,
		recipeSvc: recipeSvc,

		jobs: make(chan string, 5),
		results: make(chan RecipeError, 5),
		workerWg: workerWg,
		producerWg: producerWg,
	}
}

func (cm *CoffeeMachineService) Start() {
	go cm.result()
	go cm.createWorkerPool(3)
}

func (cm *CoffeeMachineService) MakeDrink(order []string) {
	cm.producerWg.Add(1)
	go cm.process(order)
}

func (cm *CoffeeMachineService) process(order []string) {
	for _, s := range order {
		fmt.Println("order : " + s)
		cm.jobs <- s
	}

	cm.producerWg.Done()
}

func (cm *CoffeeMachineService) worker() {
	for job := range cm.jobs {
		recipe, err := cm.recipeSvc.ByName(job)
		if err != nil {
			cm.results <- RecipeError{recipe: nil, err: err}
		} else {
			err = cm.DispenseIngredient(*recipe)
			if err != nil {
				cm.results <- RecipeError{recipe: nil, err: err}
			} else {
				cm.results <- RecipeError{recipe: recipe, err: nil}
			}
		}
	}

	cm.workerWg.Done()
}

func (cm *CoffeeMachineService) result() {
	for result := range cm.results {
		fmt.Println("Result Error is : ", result.recipe , " err : ", result.err)
		if result.err != nil {
			// call alerting service
			fmt.Println("err : ", result.err)
		} else {
			fmt.Println("serving you delicious " + result.recipe.Name)
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
	close(cm.jobs)

	cm.workerWg.Wait()
	close(cm.results)
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