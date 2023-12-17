package controllers

import (
	"product-service/source/app/products/service"

	"github.com/gin-gonic/gin"
)

type controller struct {
	service service.ServiceInterface
}

type ControllerInterface interface {
	AddProduct(ctx *gin.Context)
	ListProducts(ctx *gin.Context)
	GetProductById(ctx *gin.Context)
	UpdateProduct(ctx *gin.Context)
	DeleteProduct(ctx *gin.Context)
}
