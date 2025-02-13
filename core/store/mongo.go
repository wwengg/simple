package store

import (
	"context"
	"github.com/wwengg/simple/core/slog"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

var mongoClientInstance *mongo.Client

func MongoIns() *mongo.Client {
	if mongoClientInstance == nil {
		slog.Ins().Errorf("mongo client is nil")
		return mongoClientInstance
	}
	return mongoClientInstance
}

func NewMongoClient(mongoURI string) (*mongo.Client, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoURI).SetConnectTimeout(5*time.Second))
	if err != nil {
		slog.Ins().Errorf("mongo connect err %v", err)
		return nil, err
	}
	if err := client.Ping(context.Background(), readpref.Primary()); err != nil {
		slog.Ins().Errorf("mongo ping err %v", err)
		return nil, err
	}
	mongoClientInstance = client
	return client, nil
}

func CloseMongoClient() {
	if err := mongoClientInstance.Disconnect(context.Background()); err != nil {
		slog.Ins().Errorf("close mongo client err %v", err)
	}
}
