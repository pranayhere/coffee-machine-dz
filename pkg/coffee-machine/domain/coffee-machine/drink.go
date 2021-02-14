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

func NewBeverage(name string, recipe Recipe) *Beverage {
	return &Beverage{Name: name, Recipe: recipe}
}

func NewRecipe(name string, contents []Content) (*Recipe, error) {
	return &Recipe{Name: name, Contents: contents}, nil
}

func NewContent(ingredient Ingredient, qty int) (*Content, error) {
	return &Content{Ingredient: ingredient, Qty: qty}, nil
}

func (b *Beverage) Serve()  {
	fmt.Println(b.Name + " is prepared")
}

type RecipeRepo interface {
	Save(recipe *Recipe) error
	ByName(name string) (*Recipe, error)
}
