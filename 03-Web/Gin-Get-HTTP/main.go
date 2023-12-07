package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func sayHello(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "Hola, como estas?")
}

func sayHelloWithName(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Hola %s, como estas?", ctx.Param("name"))
}

func sayHelloWithConfindence(ctx *gin.Context) {
	query := ctx.Query("confidence")
	param := ctx.Param("name")

	if query == "si" {
		ctx.String(http.StatusOK, "Hola papa que haces")
	} else {
		ctx.String(http.StatusOK, "Hola señor %s, como le va?", param)
	}

}

func main() {
	server := gin.Default()

	sayHi := server.Group("/sayHello")

	sayHi.GET("/hi", sayHello)

	sayHi.GET("/:name", sayHelloWithName)

	sayHi.GET("/withQuery/:name", sayHelloWithConfindence)

	server.Run(":8080")
}
