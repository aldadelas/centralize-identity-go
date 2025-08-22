package mongodb_service

import (
	"context"
	"log"
	"os"
	"sync"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type MongoDBService struct {
	Client *mongo.Client
}

var (
	instance *MongoDBService
	once     sync.Once
)

func NewMongoDBService() *MongoDBService {
	defer func() {
		if err := instance.Client.Disconnect(context.Background()); err != nil {
			log.Fatal(err)
		}
	}()

	once.Do(func() {
		if os.Getenv("MONGODB_URI") == "" {
			log.Fatal("MONGODB_URI is not defined")
		}

		client, err := mongo.Connect(options.Client().ApplyURI(os.Getenv("MONGODB_URI")))
		if err != nil {
			log.Fatal(err)
		}

		instance = &MongoDBService{
			Client: client,
		}
	})

	return instance
}
