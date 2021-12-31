package database

import (
	"log"
	"time"

	"github.com/jalexanderII/zero_microservice/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitiateMongoClient() mongo.Database {
	var err error
	var client *mongo.Client
	uri := config.MONGOURI
	opts := options.Client()
	opts.ApplyURI(uri)
	opts.SetMaxPoolSize(5)
	ctx, cancel := config.NewDBContext(10 * time.Second)
	defer cancel()
	if client, err = mongo.Connect(ctx, opts); err != nil {
		log.Fatalf("Error connecting to DB: %v", err)
	}
	return *client.Database(config.CONTENTDBNAME)
}
