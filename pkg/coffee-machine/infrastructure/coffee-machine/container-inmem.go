package coffee_machine

import (
	cm "coffee-machine-dz/pkg/coffee-machine/domain/coffee-machine"
	"errors"
)

var ErrContainerNotFound = errors.New("unknown container")

type ContainerMemRepo struct {
	containers map[string]cm.Container
}

func NewContainerMemRepo() *ContainerMemRepo {
	return &ContainerMemRepo{
		containers: make(map[string]cm.Container),
	}
}

func (m *ContainerMemRepo) Save(containerToSave *cm.Container) error {
	m.containers[containerToSave.Ingredient.Name] = *containerToSave
	return nil
}

func (m *ContainerMemRepo) ByName(name string) (*cm.Container, error) {
	container, ok := m.containers[name]
	if !ok {
		return nil, ErrContainerNotFound
	}

	return &container, nil
}