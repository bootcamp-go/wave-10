package domain

// Item is a domain model that represents an item in the system.
type Item struct {
	// ID is the unique identifier of the item.
	ID    int
	// Name is the name of the item.
	Name  string
	// Price is the price of the item.
	Price float64
}