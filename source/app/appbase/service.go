package appbase

import (
	"product-service/setup"
	"product-service/source/app/products/service"
	cartService "product-service/source/app/cart/service"

)


func (b *base) WithProductService() service.ServiceInterface {
	return service.NewService(b.WithProductRepository(), setup.Logger, b.WithRedisConnection())
}

func (b *base) WithCartService() cartService.ServiceInterface {
	return cartService.NewCartService(b.WithCartRepository(), setup.Logger)
}
