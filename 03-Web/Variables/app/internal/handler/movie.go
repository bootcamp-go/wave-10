package handler

import (
	"app/internal/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

// NewMovieDefault creates a new default movie handler
func NewMovieDefault(st storage.MoviesReader) *MovieDefault {
	return &MovieDefault{st: st}
}

// MovieDefault represents the default movie handler
type MovieDefault struct {
	// st is the storage
	st storage.MoviesReader
}

// MovieJSON represents a movie entity in JSON format
type MovieJSON struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Genre string `json:"genre"`
	Year  int    `json:"year"`
}

// Get obtains all movies from the storage
func (h *MovieDefault) Get() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// request
		// ...

		// process
		movies, err := h.st.Read()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, map[string]any{"message": "error reading movies"})
			return
		}

		// response
		var data []MovieJSON
		for _, movie := range movies {
			data = append(data, MovieJSON{
				ID:    movie.ID,
				Title: movie.Title,
				Genre: movie.Genre,
				Year:  movie.Year,
			})
		}
		ctx.JSON(http.StatusOK, map[string]any{"message": "success reading movies","data": data})
	}
}