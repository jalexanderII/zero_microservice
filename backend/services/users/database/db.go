package database

import (
	"context"
	"log"
	"time"

	config "github.com/jalexanderII/zero_microservice"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToDB() mongo.Database {
	clientOptions := options.Client().ApplyURI(config.MONGOURI)
	ctx, cancel := NewDBContext(10 * time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("Error connecting to DB: %v", err)
	}
	return *client.Database(config.USERDBNAME)
}

// NewDBContext returns a new Context according to app performance
func NewDBContext(d time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), d*config.Performance/100)
}

// ConnectToTestDB overwrites DB with a Test DB
func ConnectToTestDB() mongo.Database {
	clientOptions := options.Client().ApplyURI(config.MONGOURI)
	ctx, cancel := NewDBContext(10 * time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("Error connecting to DB: %v", err)
	}
	return *client.Database(config.USERDBNAME + "_test")
}

func TestDBCleanUp(DB mongo.Database) {
	ctx, cancel := NewDBContext(30 * time.Second)
	defer cancel()
	collections, _ := DB.ListCollectionNames(ctx, bson.M{})
	for _, collection := range collections {
		ctx, cancel := NewDBContext(10 * time.Second)
		err := DB.Collection(collection).Drop(ctx)
		if err != nil {
			log.Fatalf("Error clearning DB: %v", err)
		}
		cancel()
	}
}
