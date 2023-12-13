package main

import "fmt"

type Item struct {
	ID   int
	Name string
}

type Service interface {
	// Create creates a new item
	Create(i *Item) error
	// Update updates an existing item
	Update(i *Item) error
}

// _________________________________________________________________
// in memory
type ServiceSlice struct {
	Items []Item
}

func (s *ServiceSlice) Create(i *Item) error {
	s.Items = append(s.Items, *i)
	return nil
}

func (s *ServiceSlice) Update(i *Item) error {
	for k, v := range s.Items {
		if v.ID == i.ID {
			s.Items[k] = *i
			return nil
		}
	}
	return fmt.Errorf("item with id %d not found", i.ID)
}

// _________________________________________________________________
// validator
type ServiceValidator struct {
	Service
}

func (s *ServiceValidator) Create(i *Item) error {
	// validate item
	if len(i.Name) < 5 {
		return fmt.Errorf("name must be at least 5 characters")
	}
	return s.Service.Create(i)
}

func (s *ServiceValidator) Update(i *Item) error {
	// validate item
	if len(i.Name) < 5 {
		return fmt.Errorf("name must be at least 5 characters")
	}
	return s.Service.Update(i)
}

func main() {
	// dependencies
	var sv Service
	// - in memory
	sv = &ServiceSlice{}
	// - validator
	sv = &ServiceValidator{sv}

	// create
	i := &Item{ID: 1, Name: "i1"}
	err := sv.Create(i)
	if err != nil {
		fmt.Println(err)
		return
	}
}