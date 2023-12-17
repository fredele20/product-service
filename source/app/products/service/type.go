package service

import (
	"context"
	"product-service/setup"
	"product-service/source/app/products"
	"product-service/source/app/products/models"

	"github.com/sirupsen/logrus"
)

type service struct {
	repo   products.RepositoryInterface
	redis  setup.RedisStoreInterface
	logger *logrus.Logger
}

type ServiceInterface interface {
	CreateProduct(ctx context.Context, payload *models.Product) (*models.Product, error)
	ListProducts(ctx context.Context, filter models.ListProductsParams) (*models.ListProducts, error)
	UpdateProduct(ctx context.Context, payload *models.Product) (*models.Product, error)
	DeleteProduct(ctx context.Context, id string) error
	GetProductById(ctx context.Context, id string) (*models.Product, error)
}
