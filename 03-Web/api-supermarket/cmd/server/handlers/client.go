package handlers

import (
	"api-supermarket/internal/client"

	"github.com/gin-gonic/gin"
)

type ClientRouter struct {
	group *gin.RouterGroup
	serv  client.ClientService
}

func NewClientRouter(group *gin.RouterGroup) ClientRouter {
	// creo el service y le paso el repo que cree
	serv := client.NewClientService()
	// creo un router y le paso el grupo, el service y el repo
	return ClientRouter{group, serv}
}

func (r *ClientRouter) ClientRoutes() {
	// aca van las rutas
}
