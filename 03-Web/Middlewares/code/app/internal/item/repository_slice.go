package item

import (
	"app/internal/domain"
	"fmt"
)

// NewRepositorySlice returns a new instance of RepositorySlice.
func NewRepositorySlice() *RepositorySlice {
	return &RepositorySlice{
		items: []domain.Item{},
		lastId: 0,
	}
}

// RepositorySlice is an in-memory implementation of the item.Repository interface.
type RepositorySlice struct {
	items []domain.Item
	lastId int
}

// FindAll returns all items in the repository.
func (r *RepositorySlice) FindAll() (i []domain.Item, err error) {
	// make a copy of the slice to avoid returning a reference to the internal slice
	i = make([]domain.Item, len(r.items))
	copy(i, r.items)
	return
}

// FindByID returns the item with the given ID.
func (r *RepositorySlice) FindByID(id int) (i domain.Item, err error) {
	var exists bool
	// iterate over the internal slice
	for _, item := range r.items {
		// if the item's ID matches the given ID
		if item.ID == id {
			// set the item and mark it as existing
			i = item
			exists = true
			break
		}
	}
	// check if the item exists
	if !exists {
		err = fmt.Errorf("%w - %d", ErrItemNotFound, id)
		return
	}
	
	return
}

// Save saves the given item in the repository.
func (r *RepositorySlice) Save(i *domain.Item) (err error) {
	// increment the last ID
	r.lastId++
	// set the item's ID
	i.ID = r.lastId
	// append the item to the internal slice
	r.items = append(r.items, *i)
	return
}