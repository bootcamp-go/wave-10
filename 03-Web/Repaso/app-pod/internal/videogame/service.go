package videogame

import (
	"errors"

	"github.com/LNMMusic/games-api/internal/domain"
)

var (
	// ErrServiceNotFound is an error that indicates that a videogame was not found
	ErrServiceNotFound = errors.New("service: videogame not found")
	// ErrServiceInvalid is an error that indicates that a videogame is invalid
	ErrServiceInvalid = errors.New("service: videogame invalid")
)

// Service is an interface that represents a videogame service
type Service interface {
	// Get returns a videogame by its ID
	Get(id int) (v domain.VideoGame, err error)

	// Update updates a videogame, except for the ID
	Update(v *domain.VideoGame) (err error)
}