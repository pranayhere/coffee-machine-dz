package fixture

import app "coffee-machine-dz/pkg/coffee-machine/application"

func LoadContainer(containerSvc app.ContainerService) error {
    if err := containerSvc.Save(500, 500, "hot_water"); err != nil {
        return err
    }
    if err := containerSvc.Save(500, 500, "hot_milk"); err != nil {
        return err
    }
    if err := containerSvc.Save(500, 100, "ginger_syrup"); err != nil {
        return err
    }
    if err := containerSvc.Save(500, 100, "sugar_syrup"); err != nil {
        return err
    }
    if err := containerSvc.Save(500, 100, "tea_leaves_syrup"); err != nil {
        return err
    }
    return nil
}

func LoadRecipe(recipeSvc app.RecipeService) error {
    err := recipeSvc.Save("hot_tea", map[string]int{
        "hot_water":        200,
        "hot_milk":         100,
        "ginger_syrup":     10,
        "sugar_syrup":      10,
        "tea_leaves_syrup": 30,
    })
    if err != nil {
        return err
    }

    err = recipeSvc.Save("hot_coffee", map[string]int{
        "hot_water":        100,
        "ginger_syrup":     30,
        "hot_milk":         400,
        "sugar_syrup":      50,
        "tea_leaves_syrup": 30,
    })
    if err != nil {
        return err
    }

    err = recipeSvc.Save("black_tea", map[string]int{
        "hot_water":        300,
        "ginger_syrup":     30,
        "sugar_syrup":      50,
        "tea_leaves_syrup": 30,
    })
    if err != nil {
        return err
    }

    err = recipeSvc.Save("green_tea", map[string]int{
        "hot_water":     100,
        "ginger_syrup":  30,
        "sugar_syrup":   50,
        "green_mixture": 30,
    })
    if err != nil {
        return err
    }

    return nil
}
