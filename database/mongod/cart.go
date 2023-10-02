package mongod

import (
	"context"
	"errors"
	"log"
	"product-service/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// AddToCart implements database.DataStore.
func (d *dbStore) AddToCart(ctx context.Context, userId string, productId string) error {
	var cart models.Cart
	var product models.Product
	cart.UserId = userId

	if err := d.productCollection().FindOne(ctx, bson.M{"id": productId}).Decode(&product); err != nil {
		return errors.New("Product with the given ID is not found")
	}

	if err := d.cartCollection().FindOne(ctx, bson.M{"userId": userId}).Decode(&cart); err == nil {
		_, err := d.cartCollection().UpdateOne(ctx,
			bson.D{primitive.E{Key: "userId", Value: userId}},
			bson.D{{Key: "$push", Value: bson.D{primitive.E{Key: "cartItems", Value: product}}}},
		)
		if err != nil {
			return err
		}

		return nil
	}

	cart.UserId = userId
	cart.CartItems = append(cart.CartItems, product)

	_, err := d.cartCollection().InsertOne(ctx, cart)
	if err != nil {
		return err
	}

	return nil
}

// RemoveFromCart implements database.DataStore.
func (d *dbStore) RemoveFromCart(ctx context.Context, userId string, productId string) error {
	var cart models.Cart

	err := d.cartCollection().FindOneAndUpdate(ctx,
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

// PlaceOrder implements database.DataStore.
func (d *dbStore) CheckoutCart(ctx context.Context, userId string) (*models.CartCheckoutResponse, error) {
	var cart models.Cart
	var cartCheckout models.CartCheckoutResponse
	if err := d.cartCollection().FindOne(ctx, bson.M{"userId": userId}).Decode(&cart); err != nil {
		if err.Error() == "mongo: no documents in result" {
			err = errors.New("Cart with the given userid is not found")
		}
		return nil, err
	}

	unwind := bson.D{{Key: "$unwind", Value: bson.D{primitive.E{Key: "path", Value: "$cartItems"}}}}
	group := bson.D{{Key: "$group", Value: bson.D{primitive.E{Key: "_id", Value: "$_id"}, {Key: "total", Value: bson.D{primitive.E{Key: "$sum", Value: "$cartItems.price"}}}}}}

	cursor, err := d.cartCollection().Aggregate(ctx, mongo.Pipeline{unwind, group})
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
