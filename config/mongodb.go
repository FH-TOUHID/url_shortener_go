package config
import (
	"context"
	"fmt"
	"os"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
var Client *mongo.Client
var URLCollection *mongo.Collection



func ConnectMongo() {


	uri := os.Getenv("MONGO_URI")


	client, err := mongo.Connect(
		context.TODO(),
		options.Client().ApplyURI(uri),
	)


	if err != nil {
		panic(err)
	}


	err = client.Ping(
		context.TODO(),
		nil,
	)


	if err != nil {
		panic(err)
	}


	Client = client


	fmt.Println("MongoDB connected")


	db := client.Database(
		os.Getenv("DATABASE_NAME"),
	)


	URLCollection = db.Collection(
		"urls",
	)

}