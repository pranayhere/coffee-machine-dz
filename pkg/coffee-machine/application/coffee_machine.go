package application

import (
	alerting "coffee-machine-dz/pkg/alerting/application"
	coffee_machine "coffee-machine-dz/pkg/coffee-machine/domain/coffee-machine"
	"errors"
	"sync"
)

type CoffeeMachineService struct {
	IngdSvc      IngredientSvc
	ContainerSvc ContainerSvc
	RecipeSvc    RecipeSvc
	AlertingSvc  alerting.AlertingSvc

	workerWg   sync.WaitGroup
	producerWg sync.WaitGroup

	Orders chan string // jobs
	mutex  sync.Mutex
}

type RecipeError struct {
	recipe *coffee_machine.Recipe
	err    error
}

// New
func NewCoffeeMachineService(ingdSvc IngredientSvc, containerSvc ContainerSvc, recipeSvc RecipeSvc, alertingSvc alerting.AlertingSvc) *CoffeeMachineService {
	var workerWg sync.WaitGroup
	var producerWg sync.WaitGroup
	var mutex sync.Mutex

	return &CoffeeMachineService{
		IngdSvc:      ingdSvc,
		ContainerSvc: containerSvc,
		RecipeSvc:    recipeSvc,
		AlertingSvc:  alertingSvc,

		Orders: make(chan string, 5),

		workerWg:   workerWg,
		producerWg: producerWg,

		mutex: mutex,
	}
}

type CoffeeMachineSvc interface {
	// initialize the coffee machine with number of outlets
	Init(outlets int) *CoffeeMachineService

	// fetch the recipe of the order, and create the beverage
	MakeDrink(order []string)

	// wait for completion of order and stop
	Stop()
}

func (cm *CoffeeMachineService) Init(outlets int) *CoffeeMachineService {
	cm.createWorkerPool(outlets, cm.Orders)
	return cm
}

func (cm *CoffeeMachineService) MakeDrink(order []string) {
	cm.producerWg.Add(1)
	go cm.process(order)
}

// Add order to Orders chan
func (cm *CoffeeMachineService) process(order []string) {
	for _, s := range order {
		cm.Orders <- s
	}

	cm.producerWg.Done()
}

// Process the Order and create the recipe
func (cm *CoffeeMachineService) Worker(Orders <-chan string) {
	defer cm.workerWg.Done()

	for order := range Orders {
		recipe, err := cm.RecipeSvc.ByName(order)
		if err != nil {
			cm.AlertingSvc.Alert(err)
		} else {
			r, err := cm.DispenseIngredient(*recipe)
			if err != nil {
				cm.AlertingSvc.Alert(err)
			} else {
				beverage := coffee_machine.NewBeverage(r.Name, *r)
				beverage.Serve()
			}
		}
	}
}

// Create workers equal to number of outlets
func (cm *CoffeeMachineService) createWorkerPool(noOfOutlets int, Orders <-chan string) {
	for i := 0; i < noOfOutlets; i++ {
		cm.workerWg.Add(1)
		go cm.Worker(Orders)
	}
}

// Wait for all orders to get completed and stop the machine
func (cm *CoffeeMachineService) Stop() {
	cm.producerWg.Wait()
	close(cm.Orders)

	cm.workerWg.Wait()
}

// dispense the contents of the order and return the recipe
func (cm *CoffeeMachineService) DispenseIngredient(recipe coffee_machine.Recipe) (*coffee_machine.Recipe, error) {
	cm.mutex.Lock()
	defer cm.mutex.Unlock()

	for _, content := range recipe.Contents {
		container, err := cm.ContainerSvc.ByName(content.Ingredient.Name)
		if err != nil {
			return nil, errors.New(recipe.Name + " cannot be prepared because " + content.Ingredient.Name + " is not available")
		}

		if container.Qty < content.Qty {
			return nil, errors.New(recipe.Name + " cannot be prepared because item " + container.Ingredient.Name + " is not sufficient")
		}
	}

	for _, content := range recipe.Contents {
		container, _ := cm.ContainerSvc.ByName(content.Ingredient.Name)
		_, _ = container.Dispense(content.Qty)
		err := cm.ContainerSvc.Update(container)
		if err != nil {
			return nil, err
		}
	}

	return &recipe, nil
}
