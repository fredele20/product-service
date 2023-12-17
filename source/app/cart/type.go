package cart

import (
	"context"
	"product-service/setup"
	"product-service/source/app/cart/models"
)

type cartRepository struct {
	collection setup.DBCollection
}

type RepositoryInterface interface {
	AddToCart(ctx context.Context, payload models.AddToCartRequest) error
	RemoveFromCart(ctx context.Context, userId string, productId string) error
	CheckoutCart(ctx context.Context, userId string) (*models.CartCheckoutResponse, error)
}
