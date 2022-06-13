package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"mch2022/config"
	"time"
)

var (
	client *mongo.Client
	db     *mongo.Database
)

func Connect() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var err error
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(config.MongoUrl))
	if err != nil {
		log.Fatal(err)
	}

	// Create connect
	err = client.Connect(ctx)
	if err != nil {
		log.Println("client MongoDB: " + err.Error())
	} else {
		log.Println("✔ Connected to MongoDB!")
	}

	if config.DatabaseName == "" {
		config.DatabaseName = "main"
	}

	db = client.Database(config.DatabaseName)
	if Ping() == nil {
		log.Println("✔ Pinged to MongoDB, database: " + config.DatabaseName)
		return
	}
	log.Println("can't connect to db")
	collections, err := db.ListCollections(context.Background(), nil)
	if err != nil {
		log.Println(collections)
	}
}

func Ping() error {
	return client.Ping(context.Background(), readpref.Primary())
}
