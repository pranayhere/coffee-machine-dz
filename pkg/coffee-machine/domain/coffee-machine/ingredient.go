package coffee_machine

type Ingredient struct {
	Name string
}

func NewIngredient(name string) (*Ingredient, error) {
	return &Ingredient{Name: name}, nil
}

type IngredientRepo interface {
	Save(ingredient *Ingredient) error
	ByName(name string) (*Ingredient, error)
}