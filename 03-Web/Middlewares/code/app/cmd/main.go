package main

import (
	"app/cmd/handlers"
	"app/cmd/middlewares"
	"app/internal/item"
	"fmt"
	"os"

	_ "app/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// swagger documentation
// @title Items API
// @description This is a rest api server for items.
// @version 1
// @host localhost:8080

func main() {
	// env
	token := os.Getenv("TOKEN")
	fmt.Println(token)

	// app
	// - dependencies
	rp := item.NewRepositorySlice()
	sv := item.NewServiceDefault(rp)
	hd := handlers.NewItemsDefault(sv)
	au := middlewares.NewAuthenticator(token)
	lg := middlewares.NewLogger()
	
	// - router
	rt := gin.New()
	// > middlewares
	rt.Use(lg.Log())
	rt.Use(gin.Recovery())
	// > endpoints
	// >> docs
	rt.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// >> ping
	rt.GET("/ping", func(ctx *gin.Context) {
		ctx.String(200, "pong")
	})
	// >> items group
	grItems := rt.Group("/items")
	grItems.Use(au.Authenticate())
	grItems.GET("", hd.GetItems())
	grItems.GET("/:id", hd.GetItem())
	grItems.POST("", hd.PostItem())

	// run
	if err := rt.Run(":8080"); err != nil {
		fmt.Println(err)
		return
	}
}