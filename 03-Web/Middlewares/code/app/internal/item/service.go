package item

import (
	"app/internal/domain"
	"errors"
)

var (
	// ErrServiceItemNotFound is returned when the item is not found.
	ErrServiceItemNotFound = errors.New("item not found")
)

// Service is an interface that represents a service for managing items.
type Service interface {
	// FindAll returns all items in the repository.
	FindAll() (i []domain.Item, err error)

	// FindByID returns the item with the given ID.
	FindByID(id int) (i domain.Item, err error)

	// Save saves the given item in the repository.
	Save(i *domain.Item) (err error)
}