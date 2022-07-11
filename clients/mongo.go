package clients

import (
	"context"
	"os"
	"time"

	"github.com/rjva-printerface/auth-service-go/helpers"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func NewMongoClient() *mongo.Client {
	log := helpers.NewLog("MONGO")
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URI")))

	if err != nil {
		log.Print(err.Error(), helpers.Red)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Print(err.Error(), helpers.Red)
	}
	client.Database("auth").Collection("users")

	log.Print("mongo is connected", helpers.Green)

	return client

}
