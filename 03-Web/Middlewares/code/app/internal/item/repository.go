package item

import (
	"app/internal/domain"
	"errors"
)

var (
	// ErrItemNotFound is returned when the item is not found.
	ErrItemNotFound = errors.New("repository: item not found")
)

// Repository is an interface that represents a repository of items.
type Repository interface {
	// FindAll returns all items in the repository.
	FindAll() (i []domain.Item, err error)

	// FindByID returns the item with the given ID.
	FindByID(id int) (i domain.Item, err error)

	// Save saves the given item in the repository.
	Save(i *domain.Item) (err error)
}