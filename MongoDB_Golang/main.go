package main

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

func main() {

	// get Client, Context, CancelFunc and err from connect method.
	client, ctx, cancel, err := connect("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}

	// Release resource when main function is returned.
	defer close(client, ctx, cancel)

	// Create a object of type interface to store
	// the bson values, that we are inserting into database.
	var document interface{}

	document = bson.D{
		{Key: "rollNo", Value: 175},
		{Key: "maths", Value: 80},
		{Key: "science", Value: 90},
		{Key: "computer", Value: 95},
	}

	// insertOne accepts client , context, database
	// name collection name and an interface that
	// will be inserted into the collection.
	// insertOne returns an error and a result of
	// insert in a single document into the collection.
	insertOneResult, err := insertOne(client, ctx, "gfg",
		"marks", document)

	// handle the error
	if err != nil {
		panic(err)
	}

	// print the insertion id of the document,
	// if it is inserted.
	fmt.Println("Result of InsertOne")
	fmt.Println(insertOneResult.InsertedID)

	// Now will be inserting multiple documents into
	// the collection. create a object of type slice
	// of interface to store multiple documents
	var documents []interface{}

	// Storing into interface list.
	documents = []interface{}{
		bson.D{
			{Key: "rollNo", Value: 153},
			{Key: "maths", Value: 65},
			{Key: "science", Value: 59},
			{Key: "computer", Value: 55},
		},
		bson.D{
			{Key: "rollNo", Value: 162},
			{Key: "maths", Value: 86},
			{Key: "science", Value: 80},
			{Key: "computer", Value: 69},
		},
	}

	// insertMany insert a list of documents into
	// the collection. insertMany accepts client,
	// context, database name collection name
	// and slice of interface. returns error
	// if any and result of multi document insertion.
	insertManyResult, err := insertMany(client, ctx, "gfg",
		"marks", documents)

	// handle the error
	if err != nil {
		panic(err)
	}

	fmt.Println("Result of InsertMany")

	// print the insertion ids of the multiple
	// documents, if they are inserted.
	for id := range insertManyResult.InsertedIDs {
		fmt.Println(id)
	}
}
