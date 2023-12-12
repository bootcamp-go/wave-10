package storage

import (
	"encoding/json"
	"os"
)

// NewMovieReaderJSON creates a new JSON file reader
func NewMovieReaderJSON(path string) MoviesReader {
	return &MovieReaderJSON{Path: path}
}

// MovieJSON represents a movie entity in JSON format
type MovieJSON struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Genre string `json:"genre"`
	Year  int    `json:"year"`
}

// MovieReaderJSON represents a JSON file reader
type MovieReaderJSON struct {
	// Path is the path to the JSON file
	Path string
}

// Read reads all movies from the JSON file
func (r *MovieReaderJSON) Read() (m []Movie, err error) {
	// open the file
	f, err := os.Open(r.Path)
	if err != nil {
		return
	}
	defer f.Close()

	// decode the file
	var movies []MovieJSON
	err = json.NewDecoder(f).Decode(&movies)
	if err != nil {
		return
	}

	// convert the JSON movies to the internal representation
	for _, movie := range movies {
		m = append(m, Movie{
			ID:    movie.ID,
			Title: movie.Title,
			Genre: movie.Genre,
			Year:  movie.Year,
		})
	}

	return
}