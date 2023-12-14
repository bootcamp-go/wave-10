package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/LNMMusic/games-api/internal"
	"github.com/gin-gonic/gin"
)

var (
	// ErrInvalidBodyRequest is an error that indicates that a body request is invalid
	ErrInvalidBodyRequest = errors.New("invalid body request")
)

// VideoGameJSON is a struct that represents a videogame json response
type VideoGameJSON struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Year  int    `json:"year"`
	Genre string `json:"genre"`
}

// VideoGameService is an struct that contains handlers to handle videogame requests
// for http rest-api server
type VideoGameDefault struct {
	// sv is a videogame service interface
	sv internal.VideoGameService
}

// BodyRequestVideoGamePatch is a struct that represents a videogame patch request body
type BodyRequestVideoGamePatch struct {
	Name  string `json:"name"`
	Year  int    `json:"year"`
	Genre string `json:"genre"`
}

// UpdatePartial updates a videogame, except for the ID
func (v *VideoGameDefault) UpdatePartial() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// request
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
			return
		}

		// process
		// - fetch original videogame
		videogame, err := v.sv.Get(id)
		if err != nil {
			switch err {
			case internal.ErrServiceVideoGameNotFound:
				ctx.JSON(http.StatusNotFound, gin.H{"message": "videogame not found"})
			default:
				ctx.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
			}
			return
		}
		// - serialize body request
		body := BodyRequestVideoGamePatch{
			Name:  videogame.Name,
			Year:  videogame.Year,
			Genre: videogame.Genre,
		}
		// - patch
		if err := ctx.ShouldBindJSON(&body); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": ErrInvalidBodyRequest.Error()})
			return
		}
		// - serialize videogame
		videogame.Name = body.Name
		videogame.Year = body.Year
		videogame.Genre = body.Genre
		// - update
		if err := v.sv.Update(&videogame); err != nil {
			switch err {
			case internal.ErrServiceVideoGameNotFound:
				ctx.JSON(http.StatusNotFound, gin.H{"message": "videogame not found"})
			case internal.ErrServiceVideoGameInvalid:
				ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": "invalid videogame"})
			default:
				ctx.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
			}
			return
		}

		// response
		// - serialize VideoGameJSON data
		data := VideoGameJSON{
			ID:    videogame.ID,
			Name:  videogame.Name,
			Year:  videogame.Year,
			Genre: videogame.Genre,
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "videogame updated", "data": data})
	}
}
