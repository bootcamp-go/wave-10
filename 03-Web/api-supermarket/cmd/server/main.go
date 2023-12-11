package main

import (
	"api-supermarket/cmd/server/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	// creo el server principal
	server := gin.Default()
	// creo un grupo para productos
	productsGroup := server.Group("/products")
	// llamo a la creacion de un router.
	productRouter := handlers.NewProductRouter(productsGroup)

	clientsGroup := server.Group("/clients")

	clientRouter := handlers.NewClientRouter(clientsGroup)

	// ejecuto el metodo que crea las rutas para que esten registradas
	productRouter.ProductRoutes()

	clientRouter.ClientRoutes()

	// corro mi server
	server.Run(":8080")
}
