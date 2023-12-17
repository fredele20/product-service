package appbase

import (
	"product-service/source/app/products/controllers"
	cartC "product-service/source/app/cart/controllers"


	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
)

type base struct {
	client *mongo.Client
	redisClient *redis.Client
}

type baseController struct {
	ProdC controllers.ControllerInterface
	CartC cartC.ControllerInterface
}
