package handler

import (
	"app/internal/domain"
	"app/internal/movies"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// MovieJSON represents a movie in JSON format.
type MovieJSON struct {
	ID     int     `json:"id"`
	Title  string  `json:"title"`
	Year   int     `json:"year"`
	Rating float64 `json:"rating"`
}

// BodyRequestCreateMovie represents a request body for creating a movie.
type BodyRequestCreateMovie struct {
	Title  string  `json:"title"`
	Year   int     `json:"year"`
	Rating float64 `json:"rating"`
}


// NewHandler creates a new handler with the given service layer.
func NewHandler(sv movies.Service) *Handler {
	return &Handler{
		sv: sv,
	}
}

type Handler struct {
	// sv is the service layer
	sv movies.Service
}

// AverageRating returns the average rating for all movies.
func (h *Handler) AverageRating(ctx *gin.Context) {
	// request
	// ...

	// process
	avg, err := h.sv.AverageRating()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]any{"message": "could not get the average rating"})
		return
	}

	// response
	ctx.JSON(http.StatusOK, map[string]any{
		"message": "success to get the average rating",
		"data":    avg,
	})
}

// Save saves the given movie.
func (h *Handler) Save(ctx *gin.Context) {
	// request
	var body BodyRequestCreateMovie
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]any{"message": "invalid request body"})
		return
	}

	// process
	// - deserialize the request body
	m := &domain.Movie{
		Title:  body.Title,
		Year:   body.Year,
		Rating: body.Rating,
	}
	// - save the movie
	if err := h.sv.Save(m); err != nil {
		switch {
			case errors.Is(err, movies.ErrMovieInvalidTitle) || errors.Is(err, movies.ErrMovieInvalidYear) || errors.Is(err, movies.ErrMovieInvalidRating):
				ctx.JSON(http.StatusBadRequest, map[string]any{"message": "invalid movie"})
			default:
				ctx.JSON(http.StatusInternalServerError, map[string]any{"message": "could not save the movie"})
		}
		return
	}

	// response
	ctx.JSON(http.StatusOK, map[string]any{
		"message": "success to save the movie",
		"data":    MovieJSON{
			ID:     m.ID,
			Title:  m.Title,
			Year:   m.Year,
			Rating: m.Rating,
		},
	})
}