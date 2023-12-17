package cart

import (
	"context"
	"errors"
	"log"
	"product-service/setup"
	"product-service/source/app/cart/models"
	prodModels "product-service/source/app/products/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewCartRepository(collection setup.DBCollection) RepositoryInterface {
	return &cartRepository{
		collection: collection,
	}
}

func (c *cartRepository) AddToCart(ctx context.Context, payload models.AddToCartRequest) error {
	var cart models.Cart
	var product prodModels.Product
	cart.UserId = payload.UserId

	if err := c.collection.Collection("carts").FindOne(ctx, bson.M{"id": payload.ProductId}).Decode(&product); err != nil {
		return errors.New("product with the given ID is not found")
	}

	if payload.Quantity > product.Quantity {
		err := errors.New("not enough quantity in stock, go for a lesser one")
		return err
	}

	product.Price = product.Price * uint(payload.Quantity)
	product.Quantity = payload.Quantity

	if err := c.collection.Collection("carts").FindOne(ctx, bson.M{"userId": payload.UserId}).Decode(&cart); err == nil {
		_, err := c.collection.Collection("carts").UpdateOne(ctx,
			bson.D{primitive.E{Key: "userId", Value: payload.UserId}},
			bson.D{{Key: "$push", Value: bson.D{primitive.E{Key: "cartItems", Value: product}}}},
		)
		if err != nil {
			return err
		}

		return nil
	}

	cart.CartItems = append(cart.CartItems, product)

	_, err := c.collection.Collection("carts").InsertOne(ctx, cart)
	if err != nil {
		return err
	}

	return nil
}

func (c *cartRepository) CheckoutCart(ctx context.Context, userId string) (*models.CartCheckoutResponse, error) {
	var cart models.Cart
	var cartCheckout models.CartCheckoutResponse
	if err := c.collection.Collection("carts").FindOne(ctx, bson.M{"userId": userId}).Decode(&cart); err != nil {
		if err.Error() == "mongo: no documents in result" {
			err = errors.New("cart with the given userid is not found")
		}
		return nil, err
	}

	unwind := bson.D{{Key: "$unwind", Value: bson.D{primitive.E{Key: "path", Value: "$cartItems"}}}}
	group := bson.D{{Key: "$group", Value: bson.D{primitive.E{Key: "_id", Value: "$_id"}, {Key: "total", Value: bson.D{primitive.E{Key: "$sum", Value: "$cartItems.price"}}}}}}

	cursor, err := c.collection.Collection("carts").Aggregate(ctx, mongo.Pipeline{unwind, group})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var cartItems []bson.M

	if err = cursor.All(ctx, &cartItems); err != nil {
		log.Println(err)
		return nil, err
	}

	var totalPrice int64
	for _, item := range cartItems {
		price := item["total"]
		totalPrice = price.(int64)
	}

	cartCheckout.Cart = cart
	cartCheckout.TotalPrice = int(totalPrice)

	return &cartCheckout, nil
}

func (c *cartRepository) RemoveFromCart(ctx context.Context, userId string, productId string) error {
	var cart models.Cart

	err := c.collection.Collection("carts").FindOneAndUpdate(ctx,
		bson.D{primitive.E{Key: "userId", Value: userId}},
		bson.M{"$pull": bson.M{"cartItems": bson.M{"id": productId}}},
	).Decode(&cart)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			return errors.New("document not found")
		}
		log.Println("Error: ", err)
		return err
	}

	return nil
}
