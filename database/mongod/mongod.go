package mongod

import (
	"context"
	"log"
	"product-service/database"
	"product-service/setup"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type dbStore struct {
	dbName string
	client *mongo.Client
}

var Client *mongo.Client

func MongoConnection(connectionUri, databaseName string) (database.DataStore, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionUri))
	if err != nil {
		log.Println("could not connect to mongodb...")
		log.Fatal(err)
	}

	log.Println("connected to mongodb successfully...")

	return &dbStore{client: client, dbName: databaseName}, nil
}

func (d *dbStore) productCollection() *mongo.Collection {
	return d.client.Database(d.dbName).Collection("products")
}

func (d *dbStore) cartCollection() *mongo.Collection {
	return d.client.Database(d.dbName).Collection("carts")
}

func (d *productRepository) productCollection() *mongo.Collection {
	return d.client.Database(setup.Secrets.DatabaseName).Collection("products")
}

func (d *productRepository) Collection(collectionName string) *mongo.Collection {
	return d.client.Database(setup.Secrets.DatabaseName).Collection(collectionName)
}
