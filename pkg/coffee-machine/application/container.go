package application

import cm "coffee-machine-dz/pkg/coffee-machine/domain/coffee-machine"

type ContainerService struct {
	ingredientSvc IngredientSvc

	containerRepo cm.ContainerRepo
}

func NewContainerService(ingredientSvc IngredientSvc, repo cm.ContainerRepo) *ContainerService {
	return &ContainerService{
		ingredientSvc: ingredientSvc,
		containerRepo: repo,
	}
}

type ContainerSvc interface {
	Save(cap, qty int, ingredientName string) error
	Update(container *cm.Container) error
	ByName(name string) (*cm.Container, error)
	Refill(containerName string, qty int) error
}

func (cs *ContainerService) Save(cap, qty int, ingredientName string) error {
	err := cs.ingredientSvc.Save(ingredientName)
	if err != nil {
		return err
	}

	ingd, _ := cs.ingredientSvc.ByName(ingredientName)
	container, err := cm.NewContainer(cap, qty, *ingd)
	if err != nil {
		return err
	}

	err = cs.containerRepo.Save(container)
	if err != nil {
		return err
	}

	return nil
}

func (cs *ContainerService) Update(container *cm.Container) error {
	if err := cs.containerRepo.Save(container); err != nil {
		return err
	}
	return nil
}

func (cs *ContainerService) ByName(name string) (*cm.Container, error) {
	return cs.containerRepo.ByName(name)
}

func (cs *ContainerService) Refill(containerName string, qty int) error {
	container, err := cs.ByName(containerName)
	if err != nil {
		return err
	}

	container.Qty = qty

	err = cs.containerRepo.Save(container)
	if err != nil {
		return err
	}

	return nil
}
