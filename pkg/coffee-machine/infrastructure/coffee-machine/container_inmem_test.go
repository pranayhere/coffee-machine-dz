package coffee_machine_test

import (
    domain "coffee-machine-dz/pkg/coffee-machine/domain/coffee-machine"
    infra "coffee-machine-dz/pkg/coffee-machine/infrastructure/coffee-machine"
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestContainerMemRepo(t *testing.T) {
    repo := infra.NewContainerMemRepo()

    milkContainer := addContainer(t, repo, "hot_milk")
    repoMilkContainer, err := repo.ByName("hot_milk")

    assert.NoError(t, err)
    assert.EqualValues(t, milkContainer, repoMilkContainer)

    _, err = repo.ByName("abc")
    assert.Error(t, err)
}

func addContainer(t *testing.T, repo *infra.ContainerMemRepo, ingdName string) *domain.Container {
    ingd, err := domain.NewIngredient(ingdName)
    assert.NoError(t, err)

    container, err := domain.NewContainer(500, 500, *ingd)
    assert.NoError(t, err)

    err = repo.Save(container)
    assert.NoError(t, err)

    return container
}
