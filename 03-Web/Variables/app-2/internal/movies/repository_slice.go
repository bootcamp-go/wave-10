package movies

import "app/internal/domain"

// NewRepositorySlice returns a new instance of RepositorySlice.
func NewRepositorySlice(db []domain.Movie, lastId int) *RepositorySlice {
	return &RepositorySlice{
		db: db,
		lastId: lastId,
	}
}

// RepositorySlice represents a repository for movies backed by a slice.
type RepositorySlice struct {
	db []domain.Movie
	lastId int
}

func (r *RepositorySlice) FindAll() (m []domain.Movie, err error) {
	// copy the slice to avoid returning a reference to the internal slice
	m = make([]domain.Movie, len(r.db))
	copy(m, r.db)
	return
}

func (r *RepositorySlice) Save(m *domain.Movie) (err error) {
	// consistency validations
	// ...

	// increment the lastId
	r.lastId++

	// set the ID of the movie
	m.ID = r.lastId

	// append the movie to the slice
	r.db = append(r.db, *m)
	return
}
