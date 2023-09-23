package mongod

import (
	"context"
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
	panic("unimplemented")
}

// ListProducts implements database.DataStore.
func (d *dbStore) ListProducts(ctx context.Context, payload *models.ListProductsParams) (*models.ListProducts, error) {
	panic("unimplemented")
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
