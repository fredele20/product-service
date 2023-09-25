package services

import (
	"context"
	"encoding/json"
	"product-service/models"
	"time"

	"github.com/go-redis/redis/v8"
)


func (s *Service) CreateProduct(ctx context.Context, payload *models.Product) (*models.Product, error) {
	if err := payload.Validate(); err != nil {
		s.logger.WithError(err).Error(ErrProductValidationFailed.Error())
		return nil, err
	}

	payload.Id = generateId()
	payload.TimeCreated = time.Now()
	product, err := s.db.CreateProduct(ctx, payload)
	if err != nil {
		s.logger.WithError(err).Error(ErrCouldNotAddProduct)
		return nil, ErrCouldNotAddProduct
	}
	
	return product, nil
}

func (s *Service) GetProductById(ctx context.Context, id string) (*models.Product, error) {
	product, err := s.db.GetProductById(ctx, id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (s *Service) ListProducts(ctx context.Context, payload models.ListProductsParams) (*models.ListProducts, error) {
	var result models.ListProducts

	cacheValue, err := s.redis.Get(ctx, "product_cache")
	if err == redis.Nil {
		result, err := s.db.ListProducts(ctx, payload)
		if err != nil {
			s.logger.WithError(err).Error(err.Error())
			return nil, err
		}

		cacheResult, err := json.Marshal(result)
		if err != nil {
			return nil, err
		}

		_, err = s.redis.Set(ctx, "product_cache", cacheResult, 30 * time.Second)
		if err != nil {
			return nil, err
		}

		return result, nil
	} else if err != nil {
		return nil, err
	} else {
		err = json.Unmarshal(cacheValue, &result)
		if err != nil {
			return nil, err
		}
		return &result, nil
	}
}

func (s *Service) UpdateProduct(ctx context.Context, payload *models.Product) (*models.Product, error) {
	updatedProduct, err := s.db.UpdateProduct(ctx, payload)
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