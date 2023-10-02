package mongod

import (
	"context"
	"errors"
	"fmt"
	"product-service/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// CreateProduct implements database.DataStore.
func (d *dbStore) CreateProduct(ctx context.Context, payload *models.Product) (*models.Product, error) {
	payload.ImageUri = make([]string, 0)
	if _, err := d.productCollection().InsertOne(ctx, payload); err != nil {
		return nil, err
	}
	return payload, nil
}

// DeleteProduct implements database.DataStore.
func (d *dbStore) DeleteProduct(ctx context.Context, id string) error {
	if _, err := d.productCollection().DeleteOne(ctx, bson.M{"id": id}); err != nil {
		return err
	}
	return nil
}

// GetProductById implements database.DataStore.
func (d *dbStore) GetProductById(ctx context.Context, id string) (*models.Product, error) {
	var product models.Product
	if err := d.productCollection().FindOne(ctx, bson.M{"id": id}).Decode(&product); err != nil {
		if err.Error() == "mongo: no documents in result" {
			return nil, errors.New("No product with the given id")
		}
		return nil, err
	}
	return &product, nil
}

// ListProducts implements database.DataStore.
func (d *dbStore) ListProducts(ctx context.Context) (*models.ListProducts, error) {

	filter := bson.M{}

	var products []*models.Product

	cursor, err := d.productCollection().Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	if err := cursor.All(ctx, &products); err != nil {
		fmt.Println(err)
		return nil, err
	}

	count, err := d.productCollection().CountDocuments(ctx, filter)
	if err != nil {
		return nil, err
	}

	return &models.ListProducts{
		Products: products,
		Count:    count,
	}, nil
}

// UpdateProduct implements database.DataStore.
func (d *dbStore) UpdateProduct(ctx context.Context, payload *models.Product) (*models.Product, error) {
	timeStamp := time.Now()
	payload.UpdatedAt = timeStamp
	var product models.Product
	filter := bson.M{"id": payload.Id}
	update := bson.M{"$set": payload}

	if err := d.productCollection().FindOneAndUpdate(ctx, filter, update,
		options.FindOneAndUpdate().SetReturnDocument(options.After)).Decode(&product); err != nil {
		return nil, err
	}
	return &product, nil
}

// ListProductsByStore implements database.DataStore.
func (d *dbStore) ListStoreProducts(ctx context.Context, storeId string) (*models.ListProducts, error) {
	var products []*models.Product
	
	filter := bson.M{"storeId": storeId}
	cursor, err := d.productCollection().Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	if err := cursor.All(ctx, &products); err != nil {
		return nil, err
	}

	count, err := d.productCollection().CountDocuments(ctx, filter)
	if err != nil {
		return nil, err
	}

	return &models.ListProducts{
		Products: products,
		Count: count,
	}, nil
}

