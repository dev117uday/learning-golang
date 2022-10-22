package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Person struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	FirstName string             `json:"firstname,omitempty" bson:"firstname,omitempty"`
	LastName  string             `json:"lastname,omitempty" bson:"lastname,omitempty"`
}

func main() {

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://admin:@localhost:27017/?authSource=admin"))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	defer client.Disconnect(ctx)

	dbName, _ := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Databases available : ", dbName)

	collection := client.Database("testdb").Collection("test")

	person := Person{
		FirstName: "Udayx",
		LastName:  "Yadav",
	}

	collection.InsertOne(ctx, person)

	cursor, err := collection.Find(context.TODO(), bson.D{{"firstname", "Udayx"}})
	if err != nil {
		log.Fatal(err)
	}

	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}
	for _, result := range results {
		for key, value := range result {
			fmt.Println(key, value)
		}
	}
}
