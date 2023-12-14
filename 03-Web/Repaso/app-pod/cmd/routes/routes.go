package routes

import (
	"github.com/LNMMusic/games-api/cmd/handler"
	"github.com/LNMMusic/games-api/internal/domain"
	"github.com/LNMMusic/games-api/internal/videogame"
	"github.com/gin-gonic/gin"
)

// NewRouter creates a router
func NewRouter(db map[int]domain.VideoGame, rt *gin.Engine) *Router {
	return &Router{
		db: db,
		rt: rt,
	}
}

// Router is a struct that represents a router for the server application
type Router struct {
	// db is the database of the server
	db map[int]domain.VideoGame
	// rt is the router of the server
	rt *gin.Engine
}

// MapRoutes maps the routes of the server
func (r *Router) MapRoutes() {
	// dependencies initialization
	// - repository
	rp := videogame.NewRepositoryMap(r.db)
	// - service
	sv := videogame.NewVideogameDefault(rp)
	// - handler
	hd := handler.NewVideoGameDefault(sv)

	// - router
	// > middleware
	r.rt.Use(gin.Logger())
	r.rt.Use(gin.Recovery())
	// > group /api/v1
	api := r.rt.Group("/api/v1")
	// > endpoints
	grVg := api.Group("/videogames")
	{
		grVg.PATCH("/:id", hd.UpdatePartial())
	}
}

// func (r *Router) Run() {
// 	// dependencies
// 	// - repository
// 	// - service
// 	// - handler

// 	// - router
// 	// - middleware
// 	// - endpoints

// 	// - run
// }