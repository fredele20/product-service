package service

import (
	"context"
	"encoding/json"
	"product-service/setup"
	"product-service/source/app/products"
	"product-service/source/app/products/models"
	"product-service/source/app/utils"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

func NewService(repo products.RepositoryInterface, l *logrus.Logger, redis setup.RedisStoreInterface) ServiceInterface {
	return &service{
		repo:   repo,
		logger: l,
		redis:  redis,
	}
}

func (s *service) CreateProduct(ctx context.Context, payload *models.Product) (*models.Product, error) {
	if err := payload.Validate(); err != nil {
		s.logger.WithError(err).Error(ErrProductValidationFailed.Error())
		return nil, err
	}

	payload.Id = utils.GenerateId()
	payload.TimeCreated = time.Now()
	product, err := s.repo.CreateProduct(ctx, payload)
	if err != nil {
		s.logger.WithError(err).Error(ErrCouldNotAddProduct)
		return nil, ErrCouldNotAddProduct
	}

	return product, nil
}

func (s *service) ListProducts(ctx context.Context, filter models.ListProductsParams) (*models.ListProducts, error) {
	var result models.ListProducts

	// if there is any filter arguement, fetch directly from database
	if filter.Limit != 0 || filter.StoreName != "" {
		result, err := s.repo.ListProducts(ctx, filter)
		if err != nil {
			return nil, err
		}
		return result, nil
	}

	cacheValue, err := s.redis.Get(ctx, "product_cache")
	if err == redis.Nil {
		result, err := s.repo.ListProducts(ctx, filter)
		if err != nil {
			s.logger.WithError(err).Error(err.Error())
			return nil, err
		}

		cacheResult, err := json.Marshal(result)
		if err != nil {
			return nil, err
		}

		_, err = s.redis.Set(ctx, "product_cache", cacheResult, 30*time.Second)
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

func (s *service) DeleteProduct(ctx context.Context, id string) error {
	if err := s.repo.DeleteProduct(ctx, id); err != nil {
		s.logger.WithError(err).Error(ErrCouldNotDeleteProduct)
		return err
	}
	return nil
}

func (s *service) GetProductById(ctx context.Context, id string) (*models.Product, error) {
	product, err := s.repo.GetProductById(ctx, id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (s *service) UpdateProduct(ctx context.Context, payload *models.Product) (*models.Product, error) {
	if payload.Id == "" {
		return nil, ErrInvalidProductId
	}
	updatedProduct, err := s.repo.UpdateProduct(ctx, payload)
	if err != nil {
		s.logger.WithError(err).Error(ErrCouldNotUpdateProduct)
		return nil, ErrCouldNotUpdateProduct
	}

	return updatedProduct, nil
}
