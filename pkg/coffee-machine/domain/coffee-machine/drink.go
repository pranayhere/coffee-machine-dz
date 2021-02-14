package coffee_machine

import "fmt"

type Beverage struct {
	Name   string
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

// New
func NewBeverage(name string, recipe Recipe) *Beverage {
	return &Beverage{Name: name, Recipe: recipe}
}

// New
func NewRecipe(name string, contents []Content) (*Recipe, error) {
	return &Recipe{Name: name, Contents: contents}, nil
}

// New
func NewContent(ingredient Ingredient, qty int) (*Content, error) {
	return &Content{Ingredient: ingredient, Qty: qty}, nil
}

// Serve the Beverage
func (b *Beverage) Serve() {
	fmt.Println(b.Name + " is prepared")
}

type RecipeRepo interface {
	// Save the Recipe
	Save(recipe *Recipe) error

	// Return Recipe by Name
	ByName(name string) (*Recipe, error)
}
