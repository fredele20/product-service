package services

import (
	"context"
	"product-service/models"
)

func (s *Service) AddToCart(ctx context.Context, userId, productId string) error {
	err := s.db.AddToCart(ctx, userId, productId)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) RemoveFromCart(ctx context.Context, userId, productId string) error {
	if err := s.db.RemoveFromCart(ctx, userId, productId); err != nil {
		return err
	}
	return nil
}

func (s *Service) Checkout(ctx context.Context, userId string) (*models.CartCheckoutResponse, error) {
	checkout, err := s.db.CheckoutCart(ctx, userId)
	if err != nil {
		return nil, err
	}

	return checkout, nil
}