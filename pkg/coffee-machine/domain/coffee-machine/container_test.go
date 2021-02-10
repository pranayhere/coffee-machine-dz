package coffee_machine_test

import (
	cm "coffee-machine-dz/pkg/coffee-machine/domain/coffee-machine"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewContainer(t *testing.T) {
	ingd := CreateIngredient(t)

	testContainer, err := cm.NewContainer(500, 500, *ingd)
	assert.NoError(t, err)

	assert.EqualValues(t, ingd, &testContainer.Ingredient)
	assert.EqualValues(t, 500, testContainer.Qty)
	assert.EqualValues(t, 500, testContainer.Cap)
}

func TestNewContainer_Dispense(t *testing.T) {
	ingd := CreateIngredient(t)

	testContainer, err := cm.NewContainer(500, 500, *ingd)
	assert.NoError(t, err)

	testContainer.Dispense(100)
	assert.EqualValues(t, 400, testContainer.Qty)


	_, err = testContainer.Dispense(500)
	assert.Error(t, err)
}