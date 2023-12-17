package appbase

import (
	"product-service/source/app/products/controllers"
	cartController "product-service/source/app/cart/controllers"

)

func (b *base) WithProductController() controllers.ControllerInterface {
	return controllers.NewController(b.WithProductService())
}

func (b *base) WithCartController() cartController.ControllerInterface {
	return cartController.NewController(b.WithCartService())
}
