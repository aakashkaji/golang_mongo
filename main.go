package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBsetup() *mongo.Client {

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Println("failed to connect to mongodb")
		return nil
	}
	fmt.Println("Successfully Connected to the mongodb")

	return client
}

func main() {
	// established mongo connection
	conn := DBsetup()

	// connect with projects collection
	projectCollection := conn.Database("asset_valuation").Collection("projects")

	var podcast bson.M
	// Get first record
	err := projectCollection.FindOne(context.TODO(), bson.M{}).Decode(&podcast)

	if err != nil {
		log.Fatal(err)
	}

	b, _ := json.Marshal(podcast)
	fmt.Println(string(b))
}
