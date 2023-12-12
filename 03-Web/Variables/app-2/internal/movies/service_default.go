package movies

import (
	"app/internal/domain"
	"fmt"
)

// NewServiceDefault creates a new default implementation of the Service interface.
func NewServiceDefault(rp Repository) *ServiceDefault {
	return &ServiceDefault{
		rp: rp,
	}
}

// ServiceDefault represents a default implementation of the Service interface.
type ServiceDefault struct {
	rp Repository
}

// AverageRating returns the average rating for the all movies.
func (s *ServiceDefault) AverageRating() (m float64, err error) {
	// find all movies
	mv, err := s.rp.FindAll()
	if err != nil {
		return
	}

	// calculate the average rating
	var sum float64
	for ix := range mv {
		sum += mv[ix].Rating
	}
	m = sum / float64(len(mv))
	return
}

// Save saves the given movie.
func (s *ServiceDefault) Save(m *domain.Movie) (err error) {
	// validate the movie
	// - title
	if m.Title == "" {
		err = ErrMovieInvalidTitle
		return
	}
	// - year
	if m.Year < 1900 {
		err = fmt.Errorf("%w - %d", ErrMovieInvalidYear, m.Year)
		return
	}
	// - rating
	if m.Rating < 0 || m.Rating > 10 {
		err = fmt.Errorf("%w - %f", ErrMovieInvalidRating, m.Rating)
		return
	}

	// save the movie
	err = s.rp.Save(m)
	return
}