package service

import (
	"context"
	"product-service/source/app/cart"
	"product-service/source/app/cart/models"

	"github.com/sirupsen/logrus"
)

func NewCartService(repo cart.RepositoryInterface, logger *logrus.Logger) ServiceInterface {
	return &service{
		repo: repo,
		logger: logger,
	}
}

func (s *service) AddToCart(ctx context.Context, payload models.AddToCartRequest) error {
	if err := payload.Validate(); err != nil {
		s.logger.WithError(err).Error(err.Error())
		return err
	}
	err := s.repo.AddToCart(ctx, payload)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) RemoveFromCart(ctx context.Context, userId, productId string) error {
	if err := s.repo.RemoveFromCart(ctx, userId, productId); err != nil {
		s.logger.WithError(err).Error(err.Error())
		return err
	}
	return nil
}

func (s *service) Checkout(ctx context.Context, userId string) (*models.CartCheckoutResponse, error) {
	checkout, err := s.repo.CheckoutCart(ctx, userId)
	if err != nil {
		s.logger.WithError(err).Error(err.Error())
		return nil, err
	}

	return checkout, nil
}
