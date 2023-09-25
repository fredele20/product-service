package services

import (
	"errors"
	"math/rand"
	"product-service/cache"
	"product-service/database"
	"strings"
	"time"

	"github.com/oklog/ulid"
	"github.com/sirupsen/logrus"
)

var (
	ErrCouldNotAddProduct      = errors.New("can not perform create operation right now, try again.")
	ErrCouldNotDeleteProduct   = errors.New("can not perform delete operation")
	ErrCouldNotUpdateProduct   = errors.New("can not perform update operation right now, try again.")
	ErrProductValidationFailed = errors.New("failed to validate product before persisting")
)

func generateId() string {
	entropy := ulid.Monotonic(rand.New(rand.NewSource(time.Now().Unix())), 0)
	return strings.ToLower(ulid.MustNew(ulid.Timestamp(time.Now()), entropy).String())
}

type Service struct {
	db     database.DataStore
	logger *logrus.Logger
	redis cache.CacheRedisStore
}

func NewService(db database.DataStore, l *logrus.Logger, redis cache.CacheRedisStore) *Service {
	return &Service{
		db:     db,
		logger: l,
		redis: redis,
	}
}
