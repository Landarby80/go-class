package main

import (
	"context"
	"fmt"
	"log"
	"time"

	//"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connection() {
	//URI := "mongodb://127.0.0.1:27017/?directConnection=true&serverSelectionTimeoutMS=2000&appName=mongosh+1.3.1"
	prev := time.Now()
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	_, cancle := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancle()
	if err != nil {
		log.Fatal(err)
	}
	currtime := time.Now()
	fmt.Println(currtime.Sub(prev))
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
	time.Sleep(time.Second * 10)
	// Close the connection once no longer needed
	err = client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connection to MongoDB closed.")
	}
}
