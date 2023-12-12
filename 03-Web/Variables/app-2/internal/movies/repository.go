package movies

import "app/internal/domain"

// Repository represents a repository interface for movies.
type Repository interface {
	// FindAll returns all movies stored in the repository.
	FindAll() (m []domain.Movie, err error)

	// Save stores the given movie in the repository.
	Save(m *domain.Movie) (err error)
}