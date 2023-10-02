package services

import (
	"context"
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
