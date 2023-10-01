package mongod

import (
	"context"
	"errors"
	"log"
	"product-service/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
