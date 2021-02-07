package coffee_machine

import "errors"

var ErrNotEnoughIngredient = errors.New("not enough ingredient")

type Container struct {
	Cap int
	Ingredient Ingredient
	Qty int
}

func NewContainer(cap, qty int, ingredient Ingredient) *Container {
	return &Container{
		Cap: cap,
		Qty: qty,
		Ingredient: ingredient,
	}
}

func (ic *Container) dispense(qty int) (Ingredient, error){
	if qty > ic.Qty {
		return Ingredient{}, ErrNotEnoughIngredient
	}

	ic.Qty -= qty
	return ic.Ingredient, nil
}

type ContainerRepo interface {
	Save(container *Container) error
	ByName(name string) (*Container, error)
}