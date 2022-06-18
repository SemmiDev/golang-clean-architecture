package config

import (
	"context"
	"golang-clean-architecture/exception"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoDatabase(cfg Config, env string) *mongo.Database {
	ctx, cancel := NewMongoContext()
	defer cancel()

	duration := time.Duration(cfg.MongoMaxIdleTime) * time.Second

	var option *options.ClientOptions

	if env == "test" {
		option = options.Client().
			ApplyURI(cfg.MongoURITest).
			SetMinPoolSize(cfg.MongoPoolMin).
			SetMaxPoolSize(cfg.MongoPoolMax).
			SetMaxConnIdleTime(duration)
	} else {
		option = options.Client().
			ApplyURI(cfg.MongoURI).
			SetMinPoolSize(cfg.MongoPoolMin).
			SetMaxPoolSize(cfg.MongoPoolMax).
			SetMaxConnIdleTime(duration)
	}

	client, err := mongo.NewClient(option)
	exception.PanicIfNeeded(err)

	err = client.Connect(ctx)
	exception.PanicIfNeeded(err)

	database := client.Database(cfg.MongoDB)
	return database
}

func NewMongoContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 10*time.Second)
}
