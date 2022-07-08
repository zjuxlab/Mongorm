package Mongorm

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	DbName   string
}

// var Client *mongo.Client
var Client *mongo.Client
var DB *mongo.Database

func ConnectMongo(db string, opt *options.ClientOptions) (err error) {
	Client, err := mongo.Connect(context.TODO(), opt)
	if err != nil {
		return err
	}
	err = Client.Ping(context.TODO(), nil)

	if err != nil {
		return err
	}
	DB = Client.Database(db)
	return nil
}
