package coffee_machine_test

import (
    cm "coffee-machine-dz/pkg/coffee-machine/domain/coffee-machine"
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestNewIngredient(t *testing.T) {
    testCases := []struct {
        TestName string

        Name string

        ExpectedErr bool
    }{
        {
            TestName:    "valid",
            Name:        "hot_milk",
            ExpectedErr: false,
        },
    }

    for _, c := range testCases {
        t.Run(c.TestName, func(t *testing.T) {
            ingd, err := cm.NewIngredient(c.Name)
            if c.ExpectedErr {
                assert.Error(t, err)
            } else {
                assert.NoError(t, err)
                assert.EqualValues(t, c.Name, ingd.Name)
            }
        })
    }
}

func CreateIngredient(t *testing.T) *cm.Ingredient {
    ingd, err := cm.NewIngredient("hot_milk")
    assert.NoError(t, err)

    return ingd
}
