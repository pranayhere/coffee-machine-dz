package application

import cm "coffee-machine-dz/pkg/coffee-machine/domain/coffee-machine"

type IngredientService struct {
	ingredientRepo cm.IngredientRepo
}

func NewIngredientService(repo cm.IngredientRepo) *IngredientService {
	return &IngredientService{
		ingredientRepo: repo,
	}
}

func (in *IngredientService) Save(name string) error {
	ingd, err := cm.NewIngredient(name)
	if err != nil {
		return err
	}

	err = in.ingredientRepo.Save(ingd)
	if err != nil {
		return err
	}

	return nil
}

func (in *IngredientService) ByName(name string) (*cm.Ingredient, error) {
	return in.ingredientRepo.ByName(name)
}
