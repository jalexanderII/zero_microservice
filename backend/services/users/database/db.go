package database

import (
	"log"
	"time"

	"github.com/jalexanderII/zero_microservice/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitiateMongoClient() mongo.Database {
	clientOptions := options.Client().ApplyURI(config.MONGOURI)
	ctx, cancel := config.NewDBContext(10 * time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("Error connecting to DB: %v", err)
	}
	return *client.Database(config.USERDBNAME)
}
