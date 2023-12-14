package videogame

import (
	"github.com/LNMMusic/games-api/internal/domain"
)

// NewVideogameDefault creates a videogame default service
func NewVideogameDefault(rp Repository) *ServiceDefault {
	return &ServiceDefault{
		rp: rp,
	}
}

// ServiceDefault is a struct that represents a videogame default implementation of the service
type ServiceDefault struct {
	// rp is a videogame repository
	rp Repository
}

// Get returns a videogame by its ID
func (s *ServiceDefault) Get(id int) (v domain.VideoGame, err error) {
	// get
	v, err = s.rp.Get(id)
	if err != nil {
		if err == ErrNotFound {
			err = ErrServiceNotFound
			return
		}
		return
	}

	return
}

// Update updates a videogame, except for the ID
func (s *ServiceDefault) Update(v *domain.VideoGame) (err error) {
	// validation
	// - genre
	var exists bool
	for _, genre := range domain.Genres {
		if v.Genre == genre {
			exists = true
			break
		}
	}
	if !exists {
		err = ErrServiceInvalid
		return
	}
	// - year
	if v.Year < 1958 {
		err = ErrServiceInvalid
		return
	}

	// update
	err = s.rp.Update(v)
	return
}

