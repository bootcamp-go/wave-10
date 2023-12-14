package internal

import "errors"

var (
	// ErrVideoGameNotFound is an error that indicates that a videogame was not found
	ErrVideoGameNotFound = errors.New("repository: videogame not found")
	// ErrVideoGameNameNotUnique is an error that indicates that a videogame name is not unique
	ErrVideoGameNameNotUnique = errors.New("repository: videogame name not unique")
)

// VideoGameRepository is an interface that represents a videogame repository
type VideoGameRepository interface {
	// Get returns a videogame by its ID
	Get(id int) (videogame VideoGame, err error)

	// Update updates a videogame
	Update(videogame *VideoGame) (err error)
}
