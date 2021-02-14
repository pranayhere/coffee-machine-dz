package coffee_machine

import (
	cm "coffee-machine-dz/pkg/coffee-machine/domain/coffee-machine"
	"errors"
	"sync"
)

var ErrContainerNotFound = errors.New("unknown container")

type ContainerMemRepo struct {
	containers map[string]cm.Container
	rwm        sync.RWMutex
}

func NewContainerMemRepo() *ContainerMemRepo {
	return &ContainerMemRepo{
		containers: make(map[string]cm.Container),
		rwm:        sync.RWMutex{},
	}
}

func (m *ContainerMemRepo) Save(containerToSave *cm.Container) error {
	m.rwm.Lock()
	defer m.rwm.Unlock()

	m.containers[containerToSave.Ingredient.Name] = *containerToSave
	return nil
}

func (m *ContainerMemRepo) ByName(name string) (*cm.Container, error) {
	m.rwm.RLock()
	defer m.rwm.RUnlock()

	container, ok := m.containers[name]
	if !ok {
		return nil, ErrContainerNotFound
	}

	return &container, nil
}
