package coffee_machine

import (
	cm "coffee-machine-dz/pkg/coffee-machine/domain/coffee-machine"
	"errors"
)

var ErrIngredientNotFound = errors.New("unknown ingredient")

type IngredientMemRepo struct {
	ingredients map[string]cm.Ingredient
}

func NewIngredientMemRepo() *IngredientMemRepo {
	return &IngredientMemRepo{
		ingredients: make(map[string]cm.Ingredient),
	}
}

func (m *IngredientMemRepo) Save(ingdToSave *cm.Ingredient) error {
	m.ingredients[ingdToSave.Name] = *ingdToSave
	return nil
}

func (m *IngredientMemRepo) ByName(name string) (*cm.Ingredient, error) {
	ingd, ok := m.ingredients[name]
	if !ok {
		return nil, ErrIngredientNotFound
	}

	return &ingd, nil
}