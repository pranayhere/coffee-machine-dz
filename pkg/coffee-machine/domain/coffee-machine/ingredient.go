package coffee_machine

type Ingredient struct {
	Name string
}

func NewIngredient(name string) *Ingredient {
	return &Ingredient{Name: name}
}

type IngredientRepo interface {
	Save(ingredient *Ingredient) error
	ByName(name string) (*Ingredient, error)
}