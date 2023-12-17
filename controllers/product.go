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

	payload.StoreName = ctx.GetString("storeName")

	product, err := c.services.CreateProduct(ctx, &payload)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, product)
}

// func (c *Controllers) GetProductById(ctx *gin.Context) {
// 	id := ctx.Param("id")
// 	product, err := c.services.GetProductById(ctx, id)
// 	if err != nil {
// 		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, product)
// }

// func (c *Controllers) ListProducts(ctx *gin.Context) {

// 	var filter models.ListProductsParams
// 	filter.Limit, _ = strconv.Atoi(ctx.Query("limit"))
// 	filter.StoreName = ctx.Query("storeName")

// 	products, err := c.services.ListProducts(ctx, filter)
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, products)
// }

// func (c *Controllers) UpdateProduct(ctx *gin.Context) {
// 	var payload models.Product
// 	if err := ctx.BindJSON(&payload); err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	payload.Id = ctx.Param("id")

// 	updated, err := c.services.UpdateProduct(ctx, &payload)
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, updated)
// }

// func (c *Controllers) DeleteProduct(ctx *gin.Context) {
// 	id := ctx.Param("id")

// 	if err := c.services.DeleteProduct(ctx, id); err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, "Successfully deleted the product")
// }
