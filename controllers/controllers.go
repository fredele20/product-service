package controllers

import (
	"product-service/services"

	"github.com/gin-gonic/gin"
)

type Controllers struct {
	services services.ProductServiceInterface
}

type ControllerInterface interface {
	AddProduct(ctx *gin.Context)
}

func NewController(s services.ProductServiceInterface) ControllerInterface {
	return &Controllers{
		services: s,
	}
}
