package videogame

import (
	"github.com/LNMMusic/games-api/internal/domain"
)

// NewRepositoryMap creates a videogame map repository
func NewRepositoryMap(videogames map[int]domain.VideoGame) *RepositoryMap {
	return &RepositoryMap{
		VideoGames: videogames,
	}
}

// RepositoryMap is a struct that represents a videogame map implementation of the repository
type RepositoryMap struct {
	// VideoGames is a map that represents a videogame collection
	// - key: videogame ID
	// - value: videogame
	VideoGames map[int]domain.VideoGame
}

// Get returns a videogame by its ID
func (r *RepositoryMap) Get(id int) (v domain.VideoGame, err error) {
	// check if exists
	v, ok := r.VideoGames[id]
	if !ok {
		err = ErrNotFound
		return
	}

	return
}

// Update updates a videogame
func (r *RepositoryMap) Update(v *domain.VideoGame) (err error) {
	// check if exists
	_, ok := r.VideoGames[v.ID]
	if !ok {
		err = ErrNotFound
		return
	}

	// validation
	for key := range r.VideoGames {
		// check if key differs from videogame ID and check unique name
		if key != v.ID && r.VideoGames[key].Name == v.Name {
			err = ErrNameNotUnique
			return
		}
	}
	// update
	r.VideoGames[v.ID] = *v

	return
}