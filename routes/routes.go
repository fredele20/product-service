package routes

import (
	"product-service/controllers"

	"github.com/gin-gonic/gin"
)


type Route struct {
	controller *controllers.Controllers
}

func NewRoute(c *controllers.Controllers) *Route {
	return &Route{
		controller: c,
	}
}

func Routes(incomingRoutes *gin.Engine, h Route) {
	incomingRoutes.POST("/products", h.controller.AddProduct)
	incomingRoutes.GET("/products", h.controller.ListProducts)
	incomingRoutes.GET("/products/:id", h.controller.GetProductById)
	incomingRoutes.PUT("/products/:id", h.controller.UpdateProduct)
	incomingRoutes.DELETE("/products/:id", h.controller.DeleteProduct)
}