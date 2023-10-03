package mongod

import (
	"context"
	"errors"
	"fmt"
	"product-service/models"
	"strings"
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
func (d *dbStore) ListProducts(ctx context.Context, filters models.ListProductsParams) (*models.ListProducts, error) {

	opts := options.Find()

	limit := int64(filters.Limit)

	if filters.Limit != 0 {
		opts.SetLimit(limit)
	}

	filter := bson.M{}

	if filters.StoreName != "" {
		filter["storeName"] = bson.M{"$regex": strings.TrimSpace(strings.ToLower(filters.StoreName))}
	}

	var products []*models.Product

	cursor, err := d.productCollection().Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}

	if err := cursor.All(ctx, &products); err != nil {
		fmt.Println(err)
		return nil, err
	}

	countOpts := options.Count()
	if filters.Limit != 0 {
		countOpts.SetLimit(limit)
	}
	count, err := d.productCollection().CountDocuments(ctx, filter, countOpts)
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
