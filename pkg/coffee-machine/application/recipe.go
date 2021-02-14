package application

import cm "coffee-machine-dz/pkg/coffee-machine/domain/coffee-machine"

type RecipeService struct {
    ingredientSvc IngredientSvc

    recipeRepo cm.RecipeRepo
}

// New
func NewRecipeService(ingredientSvc IngredientSvc, repo cm.RecipeRepo) *RecipeService {
    return &RecipeService{
        ingredientSvc: ingredientSvc,
        recipeRepo:    repo,
    }
}

type RecipeSvc interface {
    // Save the Recipe
    Save(name string, ingredients map[string]int) error

    // Return the Recipe by name
    ByName(name string) (*cm.Recipe, error)
}

func (rs *RecipeService) Save(name string, ingredients map[string]int) error {
    contents := make([]cm.Content, 0)

    for name, qty := range ingredients {
        err := rs.ingredientSvc.Save(name)
        if err != nil {
            return err
        }

        ingd, _ := rs.ingredientSvc.ByName(name)
        content, err := cm.NewContent(*ingd, qty)
        if err != nil {
            return err
        }

        contents = append(contents, *content)
    }

    recipe, err := cm.NewRecipe(name, contents)
    if err != nil {
        return err
    }

    err = rs.recipeRepo.Save(recipe)
    if err != nil {
        return err
    }

    return nil
}

func (rs *RecipeService) ByName(name string) (*cm.Recipe, error) {
    return rs.recipeRepo.ByName(name)
}
