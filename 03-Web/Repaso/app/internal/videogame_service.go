package internal

import "errors"

var (
	// ErrServiceVideoGameNotFound is an error that indicates that a videogame was not found
	ErrServiceVideoGameNotFound = errors.New("service: videogame not found")
	// ErrServiceVideoGameInvalid is an error that indicates that a videogame is invalid
	ErrServiceVideoGameInvalid = errors.New("service: videogame invalid")
)

// VideoGameService is an interface that represents a videogame service
type VideoGameService interface {
	// Get returns a videogame by its ID
	Get(id int) (videogame VideoGame, err error)

	// Update updates a videogame, except for the ID
	Update(videogame *VideoGame) (err error)
}