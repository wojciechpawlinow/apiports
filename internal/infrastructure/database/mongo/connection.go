package mongo

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connect is the place where actual connection with the database happens
func Connect() (*mongo.Client, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://root:secret@database:27017")) // that should be also loaded from config
	if err != nil {
		return nil, fmt.Errorf("database connection error: %w", err)
	}
	if client == nil {
		return nil, fmt.Errorf("could not connect to mongo") // errors could be changed to vars to be easily catchable
	}

	// troubles with Ping() because of missing replica?

	log.Println(fmt.Sprintf("connection with database has been established"))

	return client, nil
}
