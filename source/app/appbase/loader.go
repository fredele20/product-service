package appbase

import (
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
)


func New(client *mongo.Client, redisClient *redis.Client) *base {
	return &base{
		client: client,
		redisClient: redisClient,
	}
}

func (b *base) LoadControllers() baseController {
	var c baseController

	c.ProdC = b.WithProductController()
	c.CartC = b.WithCartController()

	return c
}
