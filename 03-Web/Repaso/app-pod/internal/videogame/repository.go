package videogame

import (
	"errors"

	"github.com/LNMMusic/games-api/internal/domain"
)

var (
	// ErrNotFound is an error that indicates that a videogame was not found
	ErrNotFound = errors.New("repository: videogame not found")
	// ErrNameNotUnique is an error that indicates that a videogame name is not unique
	ErrNameNotUnique = errors.New("repository: videogame name not unique")
)

// Repository is an interface that represents a videogame repository
type Repository interface {
	// Get returns a videogame by its ID
	Get(id int) (v domain.VideoGame, err error)

	// Update updates a videogame
	Update(v *domain.VideoGame) (err error)
}
