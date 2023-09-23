package database

import (
	"context"
	"product-service/models"
)

type DataStore interface {
	CreateProduct(ctx context.Context, payload *models.Product) (*models.Product, error)
	GetProductById(ctx context.Context, id string) (*models.Product, error)
	UpdateProduct(ctx context.Context, payload *models.Product) (*models.Product, error)
	DeleteProduct(ctx context.Context, id string) error
	ListProducts(ctx context.Context, payload *models.ListProductsParams) (*models.ListProducts, error) 
}