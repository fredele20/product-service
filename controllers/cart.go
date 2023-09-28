package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (controller *Controllers) AddToCart(ctx *gin.Context) {
	productId := ctx.Query("id")
	if productId == "" {
		log.Println("product id is empty")
		ctx.JSON(http.StatusBadRequest, gin.H{"err": "product id is empty"})
		return
	}

	userId := ctx.GetString("userid")

	cartItem, err := controller.services.AddToCart(ctx, userId, productId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "failed to add cart item"})
		return
	}

	ctx.JSON(http.StatusOK, cartItem)
}

func (controller *Controllers) RemoveFromCart(ctx *gin.Context) {
	productId := ctx.Query("id")
	if productId == "" {
		log.Println("product id is empty")
		ctx.JSON(http.StatusBadRequest, gin.H{"err": "product id is empty"})
		return
	}

	userId := ctx.GetString("userid")

	err := controller.services.RemoveFromCart(ctx, userId, productId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "failed to add cart item"})
		return
	}

	ctx.JSON(http.StatusOK, "Successfully removed item from cart")
}
