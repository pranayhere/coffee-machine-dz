package application

import (
	alerting "coffee-machine-dz/pkg/alerting/application"
	coffee_machine "coffee-machine-dz/pkg/coffee-machine/domain/coffee-machine"
	"errors"
	"sync"
)

type CoffeeMachineService struct {
	ingdSvc      IngredientSvc
	containerSvc ContainerSvc
	recipeSvc    RecipeSvc
	alertingSvc  alerting.AlertingSvc

	Orders  chan string      // jobs
	Recipes chan RecipeError // results

	workerWg   sync.WaitGroup
	producerWg sync.WaitGroup
	resultWg   sync.WaitGroup

	mutex      sync.Mutex
}

type RecipeError struct {
	recipe *coffee_machine.Recipe
	err    error
}

func NewCoffeeMachineService(ingdSvc IngredientSvc, containerSvc ContainerSvc, recipeSvc RecipeSvc, alertingSvc alerting.AlertingSvc) *CoffeeMachineService {
	var workerWg sync.WaitGroup
	var producerWg sync.WaitGroup
	var resultWg sync.WaitGroup
	var mutex sync.Mutex

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

		mutex: mutex,
	}
}

type CoffeeMachineSvc interface {
	Init(outlets int)
	MakeDrink(order []string)
	Stop()
}

func (cm *CoffeeMachineService) Init(outlets int) {
	cm.resultWg.Add(1)
	go cm.result()
	cm.createWorkerPool(outlets)
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
			r, err := cm.DispenseIngredient(*recipe)
			if err != nil {
				cm.Recipes <- RecipeError{recipe: nil, err: err}
			} else {
				cm.Recipes <- RecipeError{recipe: r, err: nil}
			}
		}
	}

	cm.workerWg.Done()
}

func (cm *CoffeeMachineService) result() {
	defer cm.resultWg.Done()

	for r := range cm.Recipes {
		if r.err != nil {
			cm.alertingSvc.Alert(r.err)
		} else {
			beverage := coffee_machine.NewBeverage(r.recipe.Name, *r.recipe)
			beverage.Serve()
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

func (cm *CoffeeMachineService) DispenseIngredient(recipe coffee_machine.Recipe) (*coffee_machine.Recipe, error) {
	cm.mutex.Lock()
	defer cm.mutex.Unlock()

	for _, content := range recipe.Contents {
		container, err := cm.containerSvc.ByName(content.Ingredient.Name)
		if err != nil {
			return nil, errors.New(recipe.Name + " cannot be prepared because " + content.Ingredient.Name + " is not available")
		}

		if container.Qty < content.Qty {
			return nil, errors.New(recipe.Name + " cannot be prepared because item " + container.Ingredient.Name + " is not sufficient")
		}
	}

	for _, content := range recipe.Contents {
		container, _ := cm.containerSvc.ByName(content.Ingredient.Name)
		_, _ = container.Dispense(content.Qty)
		err := cm.containerSvc.Update(container)
		if err != nil {
			return nil, err
		}
	}

	return &recipe, nil
}
