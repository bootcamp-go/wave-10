package item

import (
	"app/internal/domain"
	"errors"
)

// NewServiceDefault returns a new instance of ServiceDefault.
func NewServiceDefault(rp Repository) *ServiceDefault {
	return &ServiceDefault{
		rp: rp,
	}
}

// ServiceDefault is the default implementation of the item.Service interface.
type ServiceDefault struct {
	// rp is the repository used by the service.
	rp Repository
}

// FindAll returns all items in the repository.
func (s *ServiceDefault) FindAll() (i []domain.Item, err error) {
	// get all items from the repository
	i, err = s.rp.FindAll()
	if err != nil {
		return
	}
	return
}

// FindByID returns the item with the given ID.
func (s *ServiceDefault) FindByID(id int) (i domain.Item, err error) {
	// get the item from the repository
	i, err = s.rp.FindByID(id)
	if err != nil {
		return
	}
	return
}

// Save saves the given item in the repository.
func (s *ServiceDefault) Save(i *domain.Item) (err error) {
	err = s.rp.Save(i)
	if err != nil {
		if errors.Is(err, ErrItemNotFound) {
			err = ErrServiceItemNotFound
			return
		}
		return
	}
	return
}