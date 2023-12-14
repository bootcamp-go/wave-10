package repository

import (
	"sync"

	"github.com/LNMMusic/games-api/internal"
)

// NewVideoGameMap creates a videogame map repository
func NewVideoGameMap() *VideoGameMap {
	return &VideoGameMap{
		VideoGames: make(map[int]internal.VideoGame),
	}
}

// VideoGameMap is a struct that represents a videogame map implementation of the repository
type VideoGameMap struct {
	// VideoGames is a map that represents a videogame collection
	// - key: videogame ID
	// - value: videogame
	VideoGames map[int]internal.VideoGame
	// mutex is a mutex to lock the videogame map
	mutex *sync.Mutex
}

// Update updates a videogame
func (v *VideoGameMap) Update(videogame *internal.VideoGame) (err error) {
	// lock
	v.mutex.Lock()
	defer v.mutex.Unlock()
	// check if exists
	_, ok := v.VideoGames[videogame.ID]
	if !ok {
		err = internal.ErrVideoGameNotFound
		return
	}

	// validation
	for key := range v.VideoGames {
		// check if key differs from videogame ID and check unique name
		if key != videogame.ID && v.VideoGames[key].Name == videogame.Name {
			err = internal.ErrVideoGameNameNotUnique
			return
		}
	}
	// update
	v.VideoGames[videogame.ID] = *videogame

	return
}