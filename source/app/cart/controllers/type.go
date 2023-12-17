package controllers

import (
	"product-service/source/app/cart/service"

	"github.com/gin-gonic/gin"
)


type controller struct {
	service service.ServiceInterface
}

type ControllerInterface interface {
	AddToCart(ctx *gin.Context)
	RemoveFromCart(ctx *gin.Context)
	CheckoutCart(ctx *gin.Context)
}
