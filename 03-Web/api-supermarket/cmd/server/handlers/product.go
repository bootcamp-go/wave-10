package handlers

import (
	"api-supermarket/internal/product"
	"api-supermarket/pkg"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Una estructura que representa un router
type ProductRouter struct {
	// grupo al que pertenece este conjunto de urls
	productGroup *gin.RouterGroup
	// el service de productos
	service product.ProductService
}

// constructor del router de productos
func NewProductRouter(g *gin.RouterGroup) ProductRouter {
	// un slice que se rellena con una llamada al metodo util de carga de json
	slice := pkg.FullfilDB("../../products.json")
	// creo un repo y le paso el slice
	repo := product.NewProductRepository(slice)
	// creo el service y le paso el repo que cree
	serv := product.NewProductService(repo)
	// creo un router y le paso el grupo, el service y el repo
	return ProductRouter{g, serv}
}

// conjunto de rutas de URL
func (r *ProductRouter) ProductRoutes() {
	r.productGroup.GET("/ping", r.Ping())
	r.productGroup.GET("/getAll", r.GetAllProducts())
	r.productGroup.GET("/:id", r.GetById())
	r.productGroup.GET("/search", r.GetByPrice())
}

func (r *ProductRouter) Ping() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.String(200, "Pong")
	}
}

func (r *ProductRouter) GetAllProducts() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		data := r.service.GetAllProducts()
		ctx.JSON(http.StatusOK, data)
	}
}

func (r *ProductRouter) GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

		if err != nil {
			log.Println(err)
			ctx.AbortWithStatus(500)
			return
		}
		data := r.service.GetById(id)
		ctx.JSON(http.StatusOK, data)
	}
}

func (r *ProductRouter) GetByPrice() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		price, err := strconv.ParseFloat(ctx.Query("priceGt"), 64)
		if err != nil {
			log.Println(err)
			ctx.AbortWithStatus(500)
			return
		}
		data := r.service.GetByPrice(price)

		ctx.JSON(http.StatusOK, data)
	}
}
