package gomongo

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func Run(mongoUrl string) {
	// client setup
    client, err := mongo.NewClient(options.Client().ApplyURI(mongoUrl))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	// pinging client to check if it's still working
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	// listing databases
	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("databases:")
	for _, v := range databases {
		fmt.Printf("- %s\n", v)
	}

	// creating documents
	database := client.Database("nacionalpagtest")
	collection := database.Collection("clients")
	result, err := collection.InsertOne(ctx, bson.D{
		{Key: "name", Value: "Sponge Bob Square Pants"},
		{Key: "phone_number", Value: "812345678"},
		{Key: "email", Value: "spongebob@krusty.krab"},
	})
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("result:")
		fmt.Printf("%#v\n", result)
	}
}
