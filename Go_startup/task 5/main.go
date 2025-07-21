package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
)

type Trainer struct {
    Name string
    Age  int
    City string
}
func main() {
clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

// Connect to MongoDB
client, err := mongo.Connect(context.TODO(), clientOptions)

if err != nil {
    log.Fatal(err)
}

// Check the connection
err = client.Ping(context.TODO(), nil)

if err != nil {
    log.Fatal(err)
}

fmt.Println("Connected to MongoDB!")
collection := client.Database("firstdb").Collection("first")
// ash := Trainer{"Ash", 10, "Pallet Town"}
// insertResult, err := collection.InsertOne(context.TODO(), ash)
// if err != nil {
//     log.Fatal(err)
// }
// misty := Trainer{"Misty", 10, "Cerulean City"}
// brock := Trainer{"Brock", 15, "Pewter City"}
// // fmt.Println("Inserted a single document: ", insertResult.InsertedID)
// trainers := []interface{}{misty, brock}

// insertManyResult, err := collection.InsertMany(context.TODO(), trainers)
// if err != nil {
//     log.Fatal(err)
// }

// fmt.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs)
filter := bson.D{{"name", "Ash"}}

update := bson.D{
    {"$inc", bson.D{
        {"age", 1},
    }},
}
updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
if err != nil {
    log.Fatal(err)
}

fmt.Println("Matched:", updateResult.MatchedCount)
fmt.Println("Modified:", updateResult.ModifiedCount)

}