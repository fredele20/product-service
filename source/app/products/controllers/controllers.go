package controllers

import (
	"net/http"
	"product-service/source/app/products/models"
	"product-service/source/app/products/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func NewController(service service.ServiceInterface) ControllerInterface {
	return &controller{
		service: service,
	}
}

func (c *controller) AddProduct(ctx *gin.Context) {
	var payload models.Product
	if err := ctx.BindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	payload.StoreName = ctx.GetString("storeName")

	product, err := c.service.CreateProduct(ctx, &payload)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, product)
}

func (c *controller) ListProducts(ctx *gin.Context) {

	var filter models.ListProductsParams
	filter.Limit, _ = strconv.Atoi(ctx.Query("limit"))
	filter.StoreName = ctx.Query("storeName")

	products, err := c.service.ListProducts(ctx, filter)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, products)
}

func (c *controller) GetProductById(ctx *gin.Context) {
	id := ctx.Param("id")
	product, err := c.service.GetProductById(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, product)
}


func (c *controller) DeleteProduct(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := c.service.DeleteProduct(ctx, id); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, "Successfully deleted the product")
}


func (c *controller) UpdateProduct(ctx *gin.Context) {
	var payload models.Product
	if err := ctx.BindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	payload.Id = ctx.Param("id")

	updated, err := c.service.UpdateProduct(ctx, &payload)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, updated)
}
