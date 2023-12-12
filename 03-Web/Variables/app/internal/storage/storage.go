package storage

// Movie represents a movie entity
type Movie struct {
	// ID is the unique identifier of the movie
	ID     int
	// Title is the title of the movie
	Title  string
	// Genre is the genre of the movie
	Genre  string
	// Year is the year the movie was released
	Year   int
}

// MoviesReader represents a reader of movies
type MoviesReader interface {
	// Read reads all movies from the storage
	Read() (m []Movie, err error)
}