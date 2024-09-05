package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// // LoadMongoConfig reads the MongoDB URL from environment variables
// func LoadMongoConfig() (string, error) {
// 	mongoURL := os.Getenv("MONGO_URL")
// 	if mongoURL == "" {
// 		return "", fmt.Errorf("MONGO_URL is not set in the environment variables")
// 	}
// 	return mongoURL, nil
// }

var MongoDB *mongo.Database

// InitMongo initializes the MongoDB connection
func InitMongo() (*mongo.Client, context.CancelFunc, error) {

	var err error

	ctx := context.WithoutCancel(context.Background())
	ctx, cancel := context.WithTimeout(ctx, 100*time.Millisecond)

	mongoURL := Instance().MongoURL
	clientOptions := options.Client().ApplyURI(mongoURL)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, cancel, fmt.Errorf("failed to connect to MongoDB: %v", err)
	}

	// Ping the MongoDB server to verify the connection
	PingDB(client)

	MongoDB = client.Database(Instance().MongoDBName)

	return client, cancel, nil
}

func PingDB(client *mongo.Client) error {
	if err := client.Ping(context.Background(), nil); err != nil {
		return fmt.Errorf("failed to ping MongoDB: %v", err)
	}
	return nil
}

// DisconnectMongo disconnects the MongoDB client
func DisconnectMongo(client *mongo.Client, cancel context.CancelFunc) {
	defer cancel()

	defer func() {
		// client.Disconnect method also has deadline.
		// returns error if any,
		ctx := context.Background()
		if err := client.Disconnect(ctx); err != nil {
			log.Fatalf("Failed to disconnect from MongoDB: %v", err)
			panic(err)
		}
	}()
}

func GetCollection(collectionName string) *mongo.Collection {
	collection := MongoDB.Collection(collectionName)
	return collection
}
