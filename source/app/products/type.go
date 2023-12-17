package products

import (
	"context"
	"product-service/setup"
	"product-service/source/app/products/models"
)

type productRepository struct {
	collection setup.DBCollection
}

type RepositoryInterface interface {
	CreateProduct(ctx context.Context, payload *models.Product) (*models.Product, error)
	GetProductById(ctx context.Context, id string) (*models.Product, error)
	UpdateProduct(ctx context.Context, payload *models.Product) (*models.Product, error)
	DeleteProduct(ctx context.Context, id string) error
	ListProducts(ctx context.Context, filters models.ListProductsParams) (*models.ListProducts, error)
}
