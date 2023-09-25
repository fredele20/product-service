package services

import (
	"context"
	"product-service/mocks"
	"product-service/models"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

func TestService_CreateProduct(t *testing.T) {

	ctrl := gomock.NewController(t)

	dataStore := mocks.NewMockDataStore(ctrl)
	redis := mocks.NewMockCacheRedisStore(ctrl)
	logger := logrus.New()

	timeStamp := time.Now()

	// payload := models.Product{
	// 	Id:          generateId(),
	// 	Name:        "New Product",
	// 	Description: "New Description",
	// 	Quantity:    3,
	// 	Price:       20,
	// 	TimeCreated: timeStamp,
	// 	UpdatedAt:   timeStamp,
	// }
	var ctx context.Context

	t.Run("Create Product Success", func(t *testing.T) {

		payload1 := models.Product{
			Id:          generateId(),
			Name:        "New Product",
			Description: "New Description",
			Quantity:    3,
			Price:       20,
			TimeCreated: timeStamp,
			UpdatedAt:   timeStamp,
		}

		dataStore.EXPECT().CreateProduct(ctx, &payload1).Return(&payload1, nil)
		service := NewService(dataStore, logger, redis)
		product, err := service.CreateProduct(ctx, &payload1)

		require.NoError(t, err)
		require.NotEmpty(t, product)
		require.Equal(t, payload1.Id, product.Id)
		require.Equal(t, payload1.Name, product.Name)
		require.Equal(t, payload1.Description, product.Description)
		require.Equal(t, payload1.Price, product.Price)

		require.NotZero(t, product.TimeCreated)
		require.NotZero(t, product.UpdatedAt)
	})

	t.Skip()
	t.Run("Create Product Failure", func(t *testing.T) {
		payload2 := models.Product{
			Id:          generateId(),
			Name:        "",
			Description: "New Description",
			Quantity:    3,
			Price:       20,
			TimeCreated: timeStamp,
			UpdatedAt:   timeStamp,
		}
		dataStore.EXPECT().CreateProduct(ctx, &payload2).Return(&payload2, nil)
		service := NewService(dataStore, logger, redis)
		product, err := service.CreateProduct(ctx, &payload2)

		require.Error(t, err)
		require.Empty(t, product)
	})

}

// func TestService_UpdateProduct(t *testing.T) {

// 	ctrl := gomock.NewController(t)

// 	dataStore := mocks.NewMockDataStore(ctrl)
// 	logger := logrus.New()

// 	timeStamp := time.Now()

// 	payload := models.Product{
// 		Name:        "Updated Product",
// 		Description: "New Description",
// 		Quantity:    3,
// 		Price:       20,
// 		UpdatedAt:   timeStamp,
// 	}

// 	var ctx context.Context

// 	dataStore.EXPECT().UpdateProduct(ctx, &payload).Return(&payload, nil)
// 	service := NewService(dataStore, logger)
// 	product, err := service.UpdateProduct(ctx, &payload)

// 	require.NoError(t, err)
// 	require.NotEmpty(t, product)
// 	require.Equal(t, payload.Name, product.Name)
// 	require.Equal(t, payload.Description, product.Description)
// 	require.Equal(t, payload.Price, product.Price)

// 	require.NotZero(t, product.UpdatedAt)

// }

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
