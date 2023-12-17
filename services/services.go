package services

import (
	"context"
	"errors"
	"math/rand"
	"product-service/cache"
	"product-service/database"
	"product-service/database/mongod"
	"product-service/models"
	"strings"
	"time"

	"github.com/oklog/ulid"
	"github.com/sirupsen/logrus"
)

var (
	ErrCouldNotAddProduct      = errors.New("can not perform create operation right now, try again.")
	ErrCouldNotDeleteProduct   = errors.New("can not perform delete operation")
	ErrCouldNotUpdateProduct   = errors.New("can not perform update operation right now, try again.")
	ErrInvalidProductId        = errors.New("invalid product id is passed")
	ErrProductValidationFailed = errors.New("failed to validate product before persisting")
)

func GenerateId() string {
	entropy := ulid.Monotonic(rand.New(rand.NewSource(time.Now().Unix())), 0)
	return strings.ToLower(ulid.MustNew(ulid.Timestamp(time.Now()), entropy).String())
}

type Service struct {
	db     database.DataStore
	logger *logrus.Logger
	redis  cache.CacheRedisStore
}

func NewService(db database.DataStore, l *logrus.Logger, redis cache.CacheRedisStore) *Service {
	return &Service{
		db:     db,
		logger: l,
		redis:  redis,
	}
}

type prodService struct {
	prodRepo mongod.ProductRepoInterface
	// logger *logrus.Logger
}

func NewProductService(prodRepo mongod.ProductRepoInterface) ProductServiceInterface {
	return &prodService{
		prodRepo: prodRepo,
	}
}

type ProductServiceInterface interface {
	CreateProduct(ctx context.Context, payload *models.Product) (*models.Product, error)
}


// CreateProduct implements ProductServiceInterface.
func (p *prodService) CreateProduct(ctx context.Context, payload *models.Product) (*models.Product, error) {
	if err := payload.Validate(); err != nil {
		// p.logger.WithError(err).Error(ErrProductValidationFailed.Error())
		return nil, err
	}

	payload.Id = GenerateId()
	payload.TimeCreated = time.Now()
	product, err := p.prodRepo.CreateProduct(ctx, payload)
	if err != nil {
		// p.logger.WithError(err).Error(ErrCouldNotAddProduct)
		return nil, ErrCouldNotAddProduct
	}
	
	return product, nil
}
