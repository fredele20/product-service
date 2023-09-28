package services

import (
	"context"
	"product-service/models"
)


func (s *Service) AddToCart(ctx context.Context, userId, productId string) (*models.Product, error) {
	product, err := s.db.AddToCart(ctx, userId, productId)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (s *Service) RemoveFromCart(ctx context.Context, userId, productId string) error {
	if err := s.db.RemoveFromCart(ctx, userId, productId); err != nil {
		return err
	}
	return nil
}