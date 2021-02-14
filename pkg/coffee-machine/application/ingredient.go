package application

import cm "coffee-machine-dz/pkg/coffee-machine/domain/coffee-machine"

type IngredientService struct {
	ingredientRepo cm.IngredientRepo
}

// New
func NewIngredientService(repo cm.IngredientRepo) *IngredientService {
	return &IngredientService{
		ingredientRepo: repo,
	}
}

type IngredientSvc interface {
	// Save the Ingredient
	Save(name string) error

	// Return the Ingredient by name
	ByName(name string) (*cm.Ingredient, error)
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
