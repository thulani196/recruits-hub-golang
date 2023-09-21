package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoInstance : MongoInstance Struct
type MongoInstance struct {
	Client *mongo.Client
	DB     *mongo.Database
}

// MI : An instance of MongoInstance Struct
var DBInstance MongoInstance

func ConnectDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(os.Getenv("MONGO_URI")).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(ctx, opts)

	if err != nil {
		log.Fatal(err)
	}

	defer cancel()

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database connected!")

	DBInstance = MongoInstance{
		Client: client,
		DB:     client.Database(os.Getenv("DATABASE_NAME")),
	}
}
