package services

import (
	"context"
	"product-service/models"
	"time"
)


func (s *Service) CreateProduct(ctx context.Context, payload models.Product) (*models.Product, error) {
	if err := payload.Validate(); err != nil {
		s.logger.WithError(err).Error(ErrProductValidationFailed.Error())
		return nil, err
	}

	payload.Id = generateId()
	payload.TimeCreated = time.Now()
	product, err := s.db.CreateProduct(ctx, &payload)
	if err != nil {
		s.logger.WithError(err).Error(ErrCouldNotAddProduct)
		return nil, ErrCouldNotAddProduct
	}
	
	return product, nil
}

func (s *Service) UpdateProduct(ctx context.Context, payload models.Product) (*models.Product, error) {
	updatedProduct, err := s.db.UpdateProduct(ctx, &payload)
	if err != nil {
		s.logger.WithError(err).Error(ErrCouldNotUpdateProduct)
		return nil, ErrCouldNotUpdateProduct
	}

	return updatedProduct, nil
}

func (s *Service) DeleteProduct(ctx context.Context, id string) error {
	if err := s.db.DeleteProduct(ctx, id); err != nil {
		s.logger.WithError(err).Error(ErrCouldNotDeleteProduct)
		return err
	}
	return nil
}