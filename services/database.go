package services

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	user string
	pass string
	host string
	name string
	uri  string
}

func InitDatabase(user string, pass string, host string, name string) Database {
	uri := fmt.Sprintf("mongodb+srv://%s:%s@%s/%s?retryWrites=true", user, pass, host, name)
	return Database{user: user, pass: pass, host: host, name: name, uri: uri}
}

func (d *Database) Connect() (*mongo.Database, error) {
	// Create a MongoDB client
	client, err := mongo.NewClient(options.Client().ApplyURI(d.uri))
	if err != nil {
		return nil, err
	}

	// Connect to MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}
	fmt.Println("Connected to MongoDB!")

	return client.Database(d.name), nil
}

func (d *Database) GetCollection(db *mongo.Database, n string) *mongo.Collection {
	return db.Collection(n)
}
