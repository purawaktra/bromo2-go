package utils

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateMongoDb(dsn string) (*mongo.Client, error) {
	// create mongodb config
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(dsn).SetServerAPIOptions(serverAPI)

	// connect to instance
	client, err := mongo.Connect(context.TODO(), opts)

	// check if err
	if err != nil {
		Fatal(err, "CreateMongoDb", "cannot connect to mongodb")
		return nil, err
	}

	// check if instance able to make connection
	var result bson.M
	err = client.Database("bromo1").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result)
	if err != nil {
		Fatal(err, "CreateMongoDb", "failed to ping mongodb")
		return nil, err
	}

	// create return
	return client, nil
}
