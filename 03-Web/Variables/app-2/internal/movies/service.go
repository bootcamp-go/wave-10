package movies

import (
	"app/internal/domain"
	"errors"
)

var (
	// ErrMovieInvalidTitle is returned when the movie title is invalid.
	ErrMovieInvalidTitle = errors.New("service: invalid movie title")
	// ErrMovieInvalidYear is returned when the movie year is invalid.
	ErrMovieInvalidYear = errors.New("service: invalid movie year")
	// ErrMovieInvalidRating is returned when the movie rating is invalid.
	ErrMovieInvalidRating = errors.New("service: invalid movie rating")
)

// Service represents a service for managing movies.
type Service interface {
	// AverageRating returns the average rating for all movies.
	AverageRating() (m float64, err error)

	// Save saves the given movie.
	Save(m *domain.Movie) (err error)
}