package routes

import (
	"product-service/setup"
	"product-service/source/app/appbase"

	"github.com/gin-gonic/gin"
)

func RouteHandlers(r *gin.Engine) {
	app := appbase.New(setup.Client, setup.RedisClient).LoadControllers()

	r.POST("/products", app.ProdC.AddProduct)
	r.GET("/products", app.ProdC.ListProducts)
	r.GET("/products/:id", app.ProdC.GetProductById)
	r.PUT("/products/:id", app.ProdC.UpdateProduct)
	r.DELETE("/products/:id", app.ProdC.DeleteProduct)

	r.POST("/products/:id/cart", app.CartC.AddToCart)
	r.DELETE("/products/:id/cart", app.CartC.RemoveFromCart)
	r.POST("/checkout", app.CartC.CheckoutCart)
}

// type Route struct {
// 	controller *controllers.Controllers
// }

// func NewRoute(c *controllers.Controllers) *Route {
// 	return &Route{
// 		controller: c,
// 	}
// }

// func Routes(incomingRoutes *gin.Engine, h Route) {
// 	incomingRoutes.POST("/products", h.controller.AddProduct)
// 	incomingRoutes.GET("/products", h.controller.ListProducts)
// 	incomingRoutes.GET("/products/:id", h.controller.GetProductById)
// 	incomingRoutes.PUT("/products/:id", h.controller.UpdateProduct)
// 	incomingRoutes.DELETE("/products/:id", h.controller.DeleteProduct)
// 	incomingRoutes.POST("/products/:id/cart", h.controller.AddToCart)
// 	incomingRoutes.DELETE("/products/:id/cart", h.controller.RemoveFromCart)
// 	incomingRoutes.POST("/checkout", h.controller.Checkout)
// }
