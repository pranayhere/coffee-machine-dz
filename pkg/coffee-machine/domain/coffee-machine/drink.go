package coffee_machine

type Beverage struct {
	name string
	recipe Recipe
}

type Recipe struct {
	Name string
	Contents []Content
}

type Content struct {
	ingredient Ingredient
	qty int
}

func NewBeverage(name string, recipe Recipe) *Beverage {
	return &Beverage{name: name, recipe: recipe}
}

func NewRecipe(name string, contents []Content) *Recipe {
	return &Recipe{Name: name, Contents: contents}
}

func NewContent(ingredient Ingredient, qty int) *Content {
	return &Content{ingredient: ingredient, qty: qty}
}

type RecipeRepo interface {
	Save(recipe *Recipe) error
	ByName(name string) (*Recipe, error)
}