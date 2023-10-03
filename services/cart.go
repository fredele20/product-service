package services

import (
	"context"
	"product-service/models"
)

func (s *Service) AddToCart(ctx context.Context, payload models.AddToCartRequest) error {
	if err := payload.Validate(); err != nil {
		s.logger.WithError(err).Error(err.Error())
		return err
	}
	err := s.db.AddToCart(ctx, payload)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) RemoveFromCart(ctx context.Context, userId, productId string) error {
	if err := s.db.RemoveFromCart(ctx, userId, productId); err != nil {
		s.logger.WithError(err).Error(err.Error())
		return err
	}
	return nil
}

func (s *Service) Checkout(ctx context.Context, userId string) (*models.CartCheckoutResponse, error) {
	checkout, err := s.db.CheckoutCart(ctx, userId)
	if err != nil {
		s.logger.WithError(err).Error(err.Error())
		return nil, err
	}

	return checkout, nil
}