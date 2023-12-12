package domain

// Movie represents a movie in the domain layer
type Movie struct {
	// ID is the unique identifier of the movie
	ID int
	// Title is the title of the movie
	Title string
	// Year is the year the movie was released
	Year int
	// Rating is the rating of the movie
	Rating float64
}