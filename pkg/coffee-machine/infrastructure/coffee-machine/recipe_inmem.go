package coffee_machine

import (
    cm "coffee-machine-dz/pkg/coffee-machine/domain/coffee-machine"
    "errors"
)

var ErrRecipeNotFound = errors.New("recipe not found")

type RecipeMemRepo struct {
    recipes map[string]cm.Recipe
}

// New
func NewRecipeMemRepo() *RecipeMemRepo {
    return &RecipeMemRepo{
        recipes: make(map[string]cm.Recipe),
    }
}

// Save
func (m *RecipeMemRepo) Save(recipe *cm.Recipe) error {
    m.recipes[recipe.Name] = *recipe
    return nil
}

// Fetch by name
func (m *RecipeMemRepo) ByName(name string) (*cm.Recipe, error) {
    recipe, ok := m.recipes[name]
    if !ok {
        return nil, ErrRecipeNotFound
    }
    return &recipe, nil
}
