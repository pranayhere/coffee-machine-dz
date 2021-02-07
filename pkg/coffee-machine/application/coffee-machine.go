package application

import (
	"fmt"
	"sync"
)

type CoffeeMachineService struct {
	jobs chan string
	results chan string
	workerWg sync.WaitGroup
	producerWg sync.WaitGroup
}

func NewCoffeeMachineService() *CoffeeMachineService {
	var workerWg sync.WaitGroup
	var producerWg sync.WaitGroup
	cm := CoffeeMachineService{
		jobs: make(chan string, 5),
		results: make(chan string, 5),
		workerWg: workerWg,
		producerWg: producerWg,
	}

	return &cm
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
		fmt.Println("Processing drink : " + job)
		drink := job + "_PROCESSED"
		cm.results <- drink
	}
	cm.workerWg.Done()
}

func (cm *CoffeeMachineService) result() {
	for result := range cm.results {
		fmt.Println("serving you delicious " + result)
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