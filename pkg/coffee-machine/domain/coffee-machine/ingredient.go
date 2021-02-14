package coffee_machine

type Ingredient struct {
	Name string
}

// New
func NewIngredient(name string) (*Ingredient, error) {
	return &Ingredient{Name: name}, nil
}

type IngredientRepo interface {
	// Save the Ingredient
	Save(ingredient *Ingredient) error

	// Return the Ingredient by Name
	ByName(name string) (*Ingredient, error)
}
