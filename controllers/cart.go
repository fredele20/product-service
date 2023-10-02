package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (controller *Controllers) AddToCart(ctx *gin.Context) {
	productId := ctx.Param("id")
	if productId == "" {
		log.Println("product id is empty")
		ctx.JSON(http.StatusBadRequest, gin.H{"err": "product id is empty"})
		return
	}

	// NOTE: This will be used for production code.
	// userId := ctx.GetString("userid")
	// This is used for testing purposes
	userId := "jfdlfs09djfasjd34fdfj3gh"

	err := controller.services.AddToCart(ctx, userId, productId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "failed to add cart item"})
		return
	}

	ctx.JSON(http.StatusOK, "Successfully added an item to the cart")
}

func (controller *Controllers) RemoveFromCart(ctx *gin.Context) {
	productId := ctx.Param("id")
	if productId == "" {
		log.Println("product id is empty")
		ctx.JSON(http.StatusBadRequest, gin.H{"err": "product id is empty"})
		return
	}

	// userId := ctx.GetString("userid")
	userId := "jfdlfs09djfasjd34fdfj"

	err := controller.services.RemoveFromCart(ctx, userId, productId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, "Successfully removed item from cart")
}

func (controller *Controllers) Checkout(ctx *gin.Context) {

	userId := "jfdlfs09djfasjd34fdfj3gh"

	checkout, err := controller.services.Checkout(ctx, userId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, checkout)
}
