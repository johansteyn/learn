package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	//"go.mongodb.org/mongo-driver/v2/mongo"
	//"go.mongodb.org/mongo-driver/v2/options"
)

func main() {
	fmt.Println("Go 3rd-party Library: mongodb")
	fmt.Println()

	// Original driver code...
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	fmt.Println("Connecting to MongoDB...")
	//client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://johan:password@localhost:27017"))
	if err != nil {
		handleError("error connecting", err)
	}
	defer func() {
		fmt.Println("Disconnecting from MongoDB...")
		client.Disconnect(ctx)
	}()
	fmt.Printf("client=%v (%T)\n", client, client)

	/*
		// Version 2 driver code...
		fmt.Println("Connecting to MongoDB...")
		client, err := mongo.Connect(options.Client().ApplyURI("mongodb://localhost:27017"))
		if err != nil {
			handleError("error connecting", err)
		}
		fmt.Printf("client=%v (%T)\n", client, client)
	*/

	fmt.Println("Inserting a document...")
	collection := client.Database("mydb").Collection("mycollection")
	res, err := collection.InsertOne(ctx, map[string]string{"name": "Alice"})
	if err != nil {
		handleError("error inserting", err)
	}
	fmt.Printf("Inserted ID: %v\n", res.InsertedID)

	fmt.Println("Done.")

}
func handleError(message string, err error) {
	fmt.Print(message)
	fmt.Printf(" Error: %v\n", err)
	os.Exit(1)
}
