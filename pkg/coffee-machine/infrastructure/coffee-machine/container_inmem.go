package coffee_machine

import (
	cm "coffee-machine-dz/pkg/coffee-machine/domain/coffee-machine"
	"errors"
	"sync"
)

var ErrContainerNotFound = errors.New("not available")

type ContainerMemRepo struct {
	containers map[string]cm.Container
	rwm        sync.RWMutex
}

// New
func NewContainerMemRepo() *ContainerMemRepo {
	return &ContainerMemRepo{
		containers: make(map[string]cm.Container),
		rwm:        sync.RWMutex{},
	}
}

// Save
func (m *ContainerMemRepo) Save(containerToSave *cm.Container) error {
	m.rwm.Lock()
	defer m.rwm.Unlock()

	m.containers[containerToSave.Ingredient.Name] = *containerToSave
	return nil
}

// Fetch by name
func (m *ContainerMemRepo) ByName(name string) (*cm.Container, error) {
	m.rwm.RLock()
	defer m.rwm.RUnlock()

	container, ok := m.containers[name]
	if !ok {
		return nil, ErrContainerNotFound
	}

	return &container, nil
}
