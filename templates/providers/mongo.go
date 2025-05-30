package providers

import (
	"context"
	"fmt"
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var MongoProviderWireDI = wire.NewSet(NewMongoProvider)

type IMongoProvider interface {
	GetMongoClient() *mongo.Client
}

type mongoProvider struct {
	mongoClient *mongo.Client
}

func NewMongoProvider(configProvider IConfigProvider) IMongoProvider {
	// Create client options
	configMongo := configProvider.GetConfig().Mongo
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%d/%s?authSource=admin&readPreference=primary&appname=MongoDB", configMongo.Username, configMongo.Password,
		configMongo.Host, configMongo.Port, configMongo.Database)
	clientOptions := options.Client().ApplyURI(uri)

	// Initialize the MongoDB client
	client, err := mongo.Connect(clientOptions)
	if err != nil {
		panic(err)
	}
	if err = client.Ping(context.Background(), nil); err != nil {
		panic(err)
	}
	return &mongoProvider{
		mongoClient: client,
	}
}

func (m *mongoProvider) GetMongoClient() *mongo.Client {
	return m.mongoClient
}
