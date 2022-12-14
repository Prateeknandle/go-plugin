package database

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Client Database instance
var Client *mongo.Client

func DBinstance(ctx context.Context) {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
		log.Println("Wait")
	}

	mongouri := os.Getenv("MONGODB_URI")
	if mongouri == "" {
		log.Fatal("Error: not able to load MONGODB_URI")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongouri))
	if err != nil {
		panic(err)
	}

	log.Println("Connected to MongoDB!")

	Client = client

	<-ctx.Done()
	if err := client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
	log.Panicln("Disconnected MongoDB")

}
