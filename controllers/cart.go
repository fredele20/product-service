package controllers

import (
	"log"
	"net/http"
	"product-service/models"

	"github.com/gin-gonic/gin"
)

func (controller *Controllers) AddToCart(ctx *gin.Context) {
	var payload models.AddToCartRequest
	productId := ctx.Param("id")
	if productId == "" {
		log.Println("product id is empty")
		ctx.JSON(http.StatusBadRequest, gin.H{"err": "product id is empty"})
		return
	}

	if err := ctx.BindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// NOTE: This will be used for production code.
	// userId := ctx.GetString("userid")
	// This is used for testing purposes
	payload.UserId = "jfdlfs09djfasjd34fdfj3gh"
	payload.ProductId = productId

	err := controller.services.AddToCart(ctx, payload)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
