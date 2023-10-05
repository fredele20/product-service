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

func (s *ProductServiceSuite) SetupTest() {
	ctrl = gomock.NewController(s.T())

	logger = logrus.New()
	dataStore = mocks.NewMockDataStore(ctrl)
	redisStore = mocks.NewMockCacheRedisStore(ctrl)
	service = services.NewService(dataStore, logger, redisStore)
}

func TestProductServiceSuite(t *testing.T) {
	suite.Run(t, new(ProductServiceSuite))
}

func (s *ProductServiceSuite) TestService_CreateProduct() {

	s.Run("Create Product Success", func() {
		payload1 := models.Product{
			Id:          services.GenerateId(),
			Name:        "New Product",
			Description: "New Description",
			Quantity:    3,
			Price:       20,
			TimeCreated: timeStamp,
			UpdatedAt:   timeStamp,
		}

		dataStore.EXPECT().CreateProduct(ctx, &payload1).Return(&payload1, nil)
		product, err := service.CreateProduct(ctx, &payload1)

		s.NoError(err)
		s.NotEmpty(product)
		s.Equal(payload1.Id, product.Id)
		s.Equal(payload1.Name, product.Name)
		s.Equal(payload1.Description, product.Description)
		s.Equal(payload1.Price, product.Price)

		s.NotZero(product.TimeCreated)
		s.NotZero(product.UpdatedAt)
	})

	s.Run("Create Product Failure - Validation error", func() {
		payload2 := models.Product{}
		product, err := service.CreateProduct(ctx, &payload2)

		s.Error(err)
		s.Contains(err.Error(), "cannot be blank")
		s.Empty(product)
	})

}

func (s *ProductServiceSuite) TestService_UpdateProduct() {

	s.Run("Update Product Success", func() {
		payload1 := models.Product{
			Id:          services.GenerateId(),
			Name:        "New Product",
			Description: "New Description",
			Quantity:    3,
			Price:       20,
			TimeCreated: timeStamp,
			UpdatedAt:   timeStamp,
		}

		dataStore.EXPECT().CreateProduct(ctx, &payload1).Return(&payload1, nil)
		product, err := service.CreateProduct(ctx, &payload1)

		s.NoError(err)
		s.NotEmpty(product)
		s.Equal(payload1.Id, product.Id)
		s.Equal(payload1.Name, product.Name)
		s.Equal(payload1.Description, product.Description)
		s.Equal(payload1.Price, product.Price)

		s.NotZero(product.TimeCreated)
		s.NotZero(product.UpdatedAt)

		payload2 := models.Product{
			Id:        product.Id,
			Name:      "Updated Name",
			UpdatedAt: timeStamp,
		}

		dataStore.EXPECT().UpdateProduct(ctx, &payload2).Return(&payload2, nil)
		updatedProduct, err := service.UpdateProduct(ctx, &payload2)

		s.NoError(err)
		s.NotEmpty(updatedProduct)
		s.NotEqual(product.Name, updatedProduct.Name)
		s.NotEqual(product.TimeCreated, updatedProduct.UpdatedAt)

	})

	s.Run("Update Product Failure", func() {

		payload1 := models.Product{
			Id:          services.GenerateId(),
			Name:        "New Product",
			Description: "New Description",
			Quantity:    3,
			Price:       20,
			TimeCreated: timeStamp,
			UpdatedAt:   timeStamp,
		}

		dataStore.EXPECT().CreateProduct(ctx, &payload1).Return(&payload1, nil)
		product, err := service.CreateProduct(ctx, &payload1)

		s.NoError(err)
		s.NotEmpty(product)
		s.Equal(payload1.Id, product.Id)
		s.Equal(payload1.Name, product.Name)
		s.Equal(payload1.Description, product.Description)
		s.Equal(payload1.Price, product.Price)

		s.NotZero(product.TimeCreated)
		s.NotZero(product.UpdatedAt)


		payload3 := models.Product{
			Name:      "Updated Name",
			UpdatedAt: timeStamp,
		}
		dataStore.EXPECT().UpdateProduct(ctx, &payload3).Return(&payload3, nil)
		updateProductFailure, err := service.UpdateProduct(ctx, &payload3)

		s.Error(err)
		s.Empty(updateProductFailure)
	})

}

// func TestService_DeleteProduct(t *testing.T) {

// 	ctrl := gomock.NewController(t)

// 	dataStore := mocks.NewMockDataStore(ctrl)
// 	logger := logrus.New()

// 	var ctx context.Context

// 	dataStore.EXPECT().DeleteProduct(ctx, "1").Return(nil)
// 	service := NewService(dataStore, logger)
// 	err := service.DeleteProduct(ctx, "1")

// 	require.NoError(t, err)

// }

// func VerifyPhoneNumber(uuid string, otpCode int) error {
// 	var otpAccess models.OtpAccess
// 	var user models.User
// 	userResult := initializers.DB.Where("id=?", uuid).Find(&user)
// 	if userResult.RowsAffected == 0 {
// 		return services.ErrUserWithIdNotFound
// 	}

// 	otpResult := initializers.DB.Where("otp_code=?", otpCode).Find(&otpAccess)
// 	if otpResult.RowsAffected == 0 {
// 		return services.ErrInvalidOTP
// 	}

// 	currentTime := time.Now()
// 	if otpAccess.OtpExpireAt.Unix() < currentTime.Unix() {
// 		return services.ErrOTPExpired
// 	}

// 	if user.PhoneNumber != otpAccess.Receiver {
// 		return errors.New("user phone number does not match with otp receiver")
// 	}

// 	user.PhoneNumberVerifiedAt = time.Now()
// 	if err := initializers.DB.Save(&user).Error; err != nil {
// 		err = errors.New("Could not update field in user table")
// 		return err
// 	}

// 	return nil
// }
