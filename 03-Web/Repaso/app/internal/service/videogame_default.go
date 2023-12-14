package service

import "github.com/LNMMusic/games-api/internal"

// NewVideogameDefault creates a videogame default service
func NewVideogameDefault(rp internal.VideoGameRepository) *VideogameDefault {
	return &VideogameDefault{
		rp: rp,
	}
}

// VideoGameDefault is a struct that represents a videogame default implementation of the service
type VideogameDefault struct {
	// rp is a videogame repository
	rp internal.VideoGameRepository
}

// Update updates a videogame, except for the ID
func (v *VideogameDefault) Update(videogame *internal.VideoGame) (err error) {
	// validation
	// - genre
	var exists bool
	for _, genre := range internal.Genres {
		if genre == videogame.Genre {
			exists = true
			break
		}
	}
	if !exists {
		err = internal.ErrServiceVideoGameInvalid
		return
	}
	// - year
	if videogame.Year < 1958 {
		err = internal.ErrServiceVideoGameInvalid
		return
	}

	// update
	err = v.rp.Update(videogame)
	return
}

