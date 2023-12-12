package main

import (
	"app/cmd/server/handler"
	"app/internal/domain"
	"app/internal/movies"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	// env
	// ...

	// app
	// - dependencies
	rp := movies.NewRepositorySlice(make([]domain.Movie, 0), 0)
	sv := movies.NewServiceDefault(rp)
	hd := handler.NewHandler(sv)

	// - router
	rt := gin.New()
	rt.Use(gin.Recovery())
	rt.Use(gin.Logger())
	moviesGroup := rt.Group("/movies")
	moviesGroup.GET("/average-rating/", hd.AverageRating)
	moviesGroup.POST("/", hd.Save)

	// - run
	if err := rt.Run(":8080"); err != nil {
		fmt.Println(err)
		return
	}
}