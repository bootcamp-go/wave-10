package main

import (
	"fmt"

	"github.com/LNMMusic/games-api/cmd/routes"
	"github.com/LNMMusic/games-api/internal/domain"
	"github.com/gin-gonic/gin"
)

func main() {
	// env
	// ...

	// app
	// - db
	db := make(map[int]domain.VideoGame)
	// - router
	eng := gin.New()
	rt := routes.NewRouter(db, eng)
	// - map routes
	rt.MapRoutes()

	// run
	if err := eng.Run(":8080"); err != nil {
		fmt.Println(err)
		return
	}
}