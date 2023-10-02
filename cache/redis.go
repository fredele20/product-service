package cache

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

//go:generate mockgen -source=redis.go -destination=../mocks/CacheStore.go -package=mocks
type CacheRedisStore interface {
	Set(ctx context.Context, key string, value []byte, duration time.Duration) ([]byte, error)
	Get(ctx context.Context, key string) ([]byte, error)
}

type Cache struct {
	client *redis.Client
}

func RedisConnection(connectionUri string) (CacheRedisStore, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     connectionUri,
		Password: "",
		DB:       1,
	})

	fmt.Println(client)

	if err := client.Ping(context.TODO()).Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	fmt.Println("Redis connection established successfully...")

	return &Cache{
		client: client,
	}, nil
}

// Get implements CacheRedisStore.
func (c *Cache) Get(ctx context.Context, key string) ([]byte, error) {
	result, err := c.client.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	return []byte(result), nil
}

// Set implements CacheRedisStore.
func (c *Cache) Set(ctx context.Context, key string, value []byte, duration time.Duration) ([]byte, error) {
	result, err := c.client.Set(ctx, key, bytes.NewBuffer(value).Bytes(), duration).Result()
	if err != nil {
		return nil, err
	}

	return []byte(result), nil
}
