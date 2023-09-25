package controllers

import (
	"net/http"
	"product-service/models"

	"github.com/gin-gonic/gin"
)


func (c *Controllers) AddProduct(ctx *gin.Context) {
	var payload models.Product
	if err := ctx.BindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product, err := c.services.CreateProduct(ctx, &payload)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, product)
}

func (c *Controllers) UpdateProduct(ctx *gin.Context) {
	var payload models.Product
	if err := ctx.BindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := ctx.Param("id")
	payload.Id = id

	updated, err := c.services.UpdateProduct(ctx, &payload)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, updated)
}

func (c *Controllers) DeleteProduct(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := c.services.DeleteProduct(ctx, id); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, "Successfully deleted the product")
}