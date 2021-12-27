package database

import (
	"context"
	"time"

	config "github.com/jalexanderII/zero_microservice"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitiateMongoClient() (*mongo.Client, error) {
	var err error
	var client *mongo.Client
	uri := config.MONGOURI
	opts := options.Client()
	opts.ApplyURI(uri)
	opts.SetMaxPoolSize(5)
	ctx, cancel := NewDBContext(10 * time.Second)
	defer cancel()
	if client, err = mongo.Connect(ctx, opts); err != nil {
		return nil, err
	}
	return client, nil
}

// NewDBContext returns a new Context according to app performance
func NewDBContext(d time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), d*config.Performance/100)
}
