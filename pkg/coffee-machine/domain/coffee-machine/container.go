package coffee_machine

import "errors"

var ErrNotEnoughIngredient = errors.New("not enough ingredient")

type Container struct {
    Cap        int
    Qty        int
    Ingredient Ingredient
}

// New
func NewContainer(cap, qty int, ingredient Ingredient) (*Container, error) {
    return &Container{
        Cap:        cap,
        Qty:        qty,
        Ingredient: ingredient,
    }, nil
}

// Dispense the Ingredient
func (ic *Container) Dispense(qty int) (Ingredient, error) {
    if qty > ic.Qty {
        return Ingredient{}, ErrNotEnoughIngredient
    }

    ic.Qty -= qty
    return ic.Ingredient, nil
}

type ContainerRepo interface {
    // Save the container
    Save(container *Container) error

    // Return the container By Name
    ByName(name string) (*Container, error)
}
