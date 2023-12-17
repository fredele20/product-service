package appbase

import (
	"product-service/setup"
	"product-service/source/app/cart"
	"product-service/source/app/products"
)

func (b *base) WithDBCollection() setup.DBCollection {
	return setup.NewDbCollection(b.client)
}

func (b *base) WithRedisConnection() setup.RedisStoreInterface {
	return setup.NewRedisStore(b.redisClient)
}

func (b *base) WithProductRepository() products.RepositoryInterface {
	return products.NewProductRepository(b.WithDBCollection())
}

func (b *base) WithCartRepository() cart.RepositoryInterface {
	return cart.NewCartRepository(b.WithDBCollection())
}
