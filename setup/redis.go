package setup

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

type redisStore struct {
	client *redis.Client
}

var RedisClient *redis.Client

func RedisConnection() {
	client := redis.NewClient(&redis.Options{
		Addr:     Secrets.RedisUrl,
		Password: "",
		DB:       1,
	})

	if err := client.Ping(context.TODO()).Err(); err != nil {
		log.Fatal(err)
	}

	RedisClient = client

	fmt.Println("Redis connection established successfully...")

}

type RedisStoreInterface interface {
	Set(ctx context.Context, key string, value []byte, duration time.Duration) ([]byte, error)
	Get(ctx context.Context, key string) ([]byte, error)
}

func NewRedisStore(client *redis.Client) RedisStoreInterface {
	return &redisStore{
		client: client,
	}
}

func (r *redisStore) Get(ctx context.Context, key string) ([]byte, error) {
	result, err := r.client.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	return []byte(result), nil
}

func (r *redisStore) Set(ctx context.Context, key string, value []byte, duration time.Duration) ([]byte, error) {
	result, err := r.client.Set(ctx, key, bytes.NewBuffer(value).Bytes(), duration).Result()
	if err != nil {
		return nil, err
	}

	return []byte(result), nil
}
