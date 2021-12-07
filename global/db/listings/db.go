package listings

import (
	"context"
	"log"
	"time"

	"github.com/jalexanderII/zero_microservice/global"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	performance = 100
)

// DB holds Database Connection
var DB mongo.Database

func connectToDB(dbname string) {
	clientOptions := options.Client().ApplyURI(global.DBURI)
	ctx, cancel := NewDBContext(10 * time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("Error connecting to DB: %v", err)
	}
	defer client.Disconnect(ctx)

	DB = *client.Database(dbname)
}

func GetCollection(collectionName string, DB mongo.Database) mongo.Collection {
	return *DB.Collection("collectionName")
}

// NewDBContext returns a new Context according to app performance
func NewDBContext(d time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), d*performance/100)
}

// ConnectToTestDB overwrites DB with a Test DB
func ConnectToTestDB(dbname string) {
	clientOptions := options.Client().ApplyURI(global.DBURI)
	ctx, cancel := NewDBContext(10 * time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("Error connecting to DB: %v", err)
	}
	DB = *client.Database(dbname + "_test")

	ctx, cancel = NewDBContext(30 * time.Second)
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
