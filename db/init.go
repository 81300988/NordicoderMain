package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"nc_student.com/v1/config"
)

var Client *mongo.Client

func init() {
	connect()
}

func connect() {
	client, err := mongo.NewClient(options.Client().ApplyURI(config.Config.Mongo.Uri))
	if err != nil {
		log.Fatalf("connect error :%v", err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	ctx, _ = context.WithTimeout(context.Background(), 2*time.Second)
	err = client.Ping(ctx, readpref.Primary())

	Client = client
}
