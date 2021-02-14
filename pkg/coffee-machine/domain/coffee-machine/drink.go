package coffee_machine

type Beverage struct {
	Recipe Recipe
}

type Recipe struct {
	Name     string
	Contents []Content
}

type Content struct {
	Ingredient Ingredient
	Qty        int
}

func NewBeverage(name string, recipe Recipe) *Beverage {
	return &Beverage{Recipe: recipe}
}

func NewRecipe(name string, contents []Content) (*Recipe, error) {
	return &Recipe{Name: name, Contents: contents}, nil
}

func NewContent(ingredient Ingredient, qty int) (*Content, error) {
	return &Content{Ingredient: ingredient, Qty: qty}, nil
}

type RecipeRepo interface {
	Save(recipe *Recipe) error
	ByName(name string) (*Recipe, error)
}
