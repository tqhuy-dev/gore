package providers

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type IMongoProvider interface {
	GetMongoClient() *mongo.Client
}

type mongoProvider struct {
	mongoClient *mongo.Client
}

type IMongoConfig interface {
	GetUri() string
}

func NewMongoProvider(mongoConfig IMongoConfig) *mongo.Client {
	// Create client options
	//mongodb://%s:%s@%s:%d/%s?authSource=admin&readPreference=primary&appname=MongoDB
	clientOptions := options.Client().ApplyURI(mongoConfig.GetUri())

	// Initialize the MongoDB client
	client, err := mongo.Connect(clientOptions)
	if err != nil {
		panic(err)
	}
	if err = client.Ping(context.Background(), nil); err != nil {
		panic(err)
	}
	return client
}
