package application

import cm "coffee-machine-dz/pkg/coffee-machine/domain/coffee-machine"

type RecipeService struct {
	ingredientSvc IngredientSvc

	recipeRepo cm.RecipeRepo
}

func NewRecipeService(ingredientSvc IngredientSvc, repo cm.RecipeRepo) *RecipeService {
	return &RecipeService{
		ingredientSvc: ingredientSvc,
		recipeRepo:    repo,
	}
}

type RecipeSvc interface {
	Save(name string, ingredients map[string]int) error
	ByName(name string) (*cm.Recipe, error)
}

func (rs *RecipeService) Save(name string, ingredients map[string]int) error {
	contents := make([]cm.Content, 0)

	for name, qty := range ingredients {
		ingd, err := rs.ingredientSvc.ByName(name)
		if err != nil {
			return err
		}

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
