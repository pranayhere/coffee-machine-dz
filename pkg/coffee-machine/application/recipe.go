package application

import cm "coffee-machine-dz/pkg/coffee-machine/domain/coffee-machine"

type RecipeService struct {
	ingredientSvc IngredientService

	recipeRepo cm.RecipeRepo
}

func NewRecipeService(ingredientSvc IngredientService, repo cm.RecipeRepo) *RecipeService {
	return &RecipeService{
		ingredientSvc: ingredientSvc,
		recipeRepo:    repo,
	}
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
