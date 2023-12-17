package service

import (
	"context"
	"product-service/source/app/cart"
	"product-service/source/app/cart/models"

	"github.com/sirupsen/logrus"
)

type service struct {
	repo cart.RepositoryInterface
	logger *logrus.Logger
}

type ServiceInterface interface {
	AddToCart(ctx context.Context, payload models.AddToCartRequest) error
	RemoveFromCart(ctx context.Context, userId, productId string) error
	Checkout(ctx context.Context, userId string) (*models.CartCheckoutResponse, error)
}
