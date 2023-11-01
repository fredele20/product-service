package services_test

import (
	"context"
	"product-service/mocks"
	"product-service/models"
	"product-service/services"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
)

type ProductServiceSuite struct {
	suite.Suite
}

var ctrl *gomock.Controller
var dataStore *mocks.MockDataStore
var redisStore *mocks.MockCacheRedisStore
var logger *logrus.Logger
var service *services.Service
var timeStamp = time.Now()
var ctx context.Context
var payload models.Product
var err error

func (s *ProductServiceSuite) SetupTest() {
	ctrl = gomock.NewController(s.T())

	logger = logrus.New()
	dataStore = mocks.NewMockDataStore(ctrl)
	redisStore = mocks.NewMockCacheRedisStore(ctrl)
	service = services.NewService(dataStore, logger, redisStore)

	payload = models.Product{
		Id:          services.GenerateId(),
		Name:        "New Product",
		Description: "New Description",
		Quantity:    3,
		Price:       20,
		TimeCreated: timeStamp,
		UpdatedAt:   timeStamp,
	}
}

func TestProductServiceSuite(t *testing.T) {
	suite.Run(t, new(ProductServiceSuite))
}

func (s *ProductServiceSuite) TestService_CreateProduct() {
	// Success path
	dataStore.EXPECT().CreateProduct(ctx, &payload).Return(&payload, err)
	product, err := service.CreateProduct(ctx, &payload)

	s.NoError(err)
	s.NotEmpty(product)
	s.Equal(payload.Id, product.Id)
	s.Equal(payload.Name, product.Name)
	s.Equal(payload.Description, product.Description)
	s.Equal(payload.Price, product.Price)

	s.NotZero(product.TimeCreated)
	s.NotZero(product.UpdatedAt)

	// Failure path, due to input validation
	payload = models.Product{}
	product, err = service.CreateProduct(ctx, &payload)

	s.Error(err)
	s.Contains(err.Error(), "cannot be blank")
	s.Empty(product)

}

func (s *ProductServiceSuite) TestService_UpdateProduct() {
	dataStore.EXPECT().CreateProduct(ctx, &payload).Return(&payload, err)
	product, err := service.CreateProduct(ctx, &payload)

	s.NoError(err)
	s.NotEmpty(product)
	s.Equal(payload.Id, product.Id)
	s.NotZero(product.TimeCreated)
	s.NotZero(product.UpdatedAt)

	payload2 := product

	payload2.Name = "Updated Name"
	payload2.Description = "Updated Value"

	dataStore.EXPECT().UpdateProduct(ctx, payload2).Return(payload2, err)
	updatedProduct, err := service.UpdateProduct(ctx, payload2)

	s.NoError(err)
	s.NotEmpty(updatedProduct)
	s.Equal(product.Name, updatedProduct.Name)
	s.Equal(product.Description, updatedProduct.Description)
	s.Equal(product.UpdatedAt, updatedProduct.UpdatedAt)

	// Failure path - invalid id
	payload2.Id = ""
	updateProductFailure, err := service.UpdateProduct(ctx, payload2)

	s.Error(err)
	s.Empty(updateProductFailure)

}

func (s *ProductServiceSuite) TestService_GetProductById() {

	dataStore.EXPECT().CreateProduct(ctx, &payload).Return(&payload, err)
	product, err := service.CreateProduct(ctx, &payload)

	s.NoError(err)
	s.NotEmpty(product)

	// Success part - valid id arguement
	id := product.Id

	dataStore.EXPECT().GetProductById(ctx, id).Times(1).Return(product, nil)
	singleProduct, err := service.GetProductById(ctx, id)

	s.NoError(err)
	s.NotEmpty(singleProduct)
	s.Equal(product.Name, singleProduct.Name)
	s.Equal(product.Description, singleProduct.Description)

	// Failure part - Invalid id argument
	id = ""
	dataStore.EXPECT().GetProductById(ctx, id).Return(nil, err)
	noProduct, err := service.GetProductById(ctx, id)

	s.NoError(err)
	s.Empty(noProduct)

}

func (s *ProductServiceSuite) TestService_DeleteProduct() {
	dataStore.EXPECT().CreateProduct(ctx, &payload).Return(&payload, nil)
	product, err := service.CreateProduct(ctx, &payload)

	s.NoError(err)
	s.NotEmpty(product)

	id := product.Id

	dataStore.EXPECT().DeleteProduct(ctx, id).Return(nil)
	err = service.DeleteProduct(ctx, id)

	s.NoError(err)
}