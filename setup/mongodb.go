package setup

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

type dbCollection struct {
	client *mongo.Client
}

func ConnectMongo() {
	var ctx, cancel = context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(Secrets.DatabaseURL))
	if err != nil {
		log.Println("could not connect to mongodb...")
		log.Fatal(err)
	}
	Client = client
	log.Println("connected to mongodb successfully...")
}

type DBCollection interface {
	Collection(collectionName string) *mongo.Collection
}

func NewDbCollection(client *mongo.Client) DBCollection {
	return &dbCollection{
		client: client,
	}
}

func (d *dbCollection) Collection(collectionName string) *mongo.Collection {
	return d.client.Database(Secrets.DatabaseName).Collection(collectionName)
}
