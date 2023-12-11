package handler

import (
	"app/internal"
	"net/http"

	"github.com/gin-gonic/gin"
)

// NewDefaultMovies creates a new instance of DefaultMovies
func NewDefaultMovies() *DefaultMovies {
	return &DefaultMovies{
		db: make([]internal.Movie, 0),
		lastID: 0,
	}
}

// DefaultMovies is an struct that represents the handler of movies
type DefaultMovies struct {
	// db is an slice of movies
	db []internal.Movie
	// lastID is the last ID of the movie
	lastID int
}

// MovieJSON is an struct that represents the response of a movie in json format
type MovieJSON struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Year      int    `json:"year"`
	Director  string `json:"director"`
	InTheater bool   `json:"in_theater"`
}

// BodyRequestCreate is an struct that represents the body request of a movie in json format
type BodyRequestCreate struct {
	Title     string `json:"title"`
	Year      int    `json:"year"`
	Director  string `json:"director"`
	InTheater bool   `json:"in_theater"`
}
// CreateMovie creates a new movie
func (d *DefaultMovies) CreateMovie() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// authentication
		token := ctx.GetHeader("Authorization")
		if token != "123456" {
			ctx.JSON(http.StatusUnauthorized, map[string]any{"message": "invalid token"})
			return
		}

		// request
		var body BodyRequestCreate
		err := ctx.ShouldBindJSON(&body)
		// err := json.NewDecoder(ctx.Request.Body).Decode(&body)
		if err != nil {
			// b, _ := json.Marshal(map[string]string{
			// 	"message": "invalid request body",
			// })
			// ctx.Writer.WriteHeader(http.StatusBadRequest)
			// ctx.Writer.Write(b)
			// ctx.Writer.Header().Set("Content-Type", "application/json")
			ctx.JSON(http.StatusBadRequest, map[string]any{"message": "invalid request body"})
			return
		}

		// process
		// - deserialize
		mv := internal.Movie{
			Title:     body.Title,
			Year:      body.Year,
			Director:  body.Director,
			InTheater: body.InTheater,
		}
		// - validation (required / quality / business logic - previous to storage - agnostic from who instanciate the data)
		// if mv.Title == "" {
		// 	ctx.JSON(http.StatusBadRequest, map[string]any{"message": "title is required"})
		// 	return
		// }
		// if mv.Year < 0 {
		// 	ctx.JSON(http.StatusBadRequest, map[string]any{"message": "year is required"})
		// 	return
		// }
		d.lastID++
		mv.ID = d.lastID
		// save
		d.db = append(d.db, mv)

		// response
		ctx.JSON(http.StatusCreated, map[string]any{
			"message": "movie created",
			"data": MovieJSON{
				ID:        mv.ID,
				Title:     mv.Title,
				Year:      mv.Year,
				Director:  mv.Director,
				InTheater: mv.InTheater,
			},
		})
	}
}
