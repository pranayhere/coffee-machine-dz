package application

import cm "coffee-machine-dz/pkg/coffee-machine/domain/coffee-machine"

type ContainerService struct {
	ingredientSvc IngredientService

	containerRepo cm.ContainerRepo
}

func NewContainerService(ingredientSvc IngredientService, repo cm.ContainerRepo) *ContainerService {
	return &ContainerService{
		ingredientSvc: ingredientSvc,
		containerRepo: repo,
	}
}

func (cs *ContainerService) Save(cap, qty int, ingredientName string) error {
	ingd, err := cs.ingredientSvc.ByName(ingredientName)
	if err != nil {
		return err
	}

	err = cs.containerRepo.Save(cm.NewContainer(cap, qty, *ingd))
	if err != nil {
		return err
	}

	return nil
}

func (cs *ContainerService) ByName(name string) (*cm.Container, error) {
	return cs.containerRepo.ByName(name)
}