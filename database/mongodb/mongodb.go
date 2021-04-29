package mongodb

import (
	"context"
	"github.com/maiaaraujo5/gostart/database"
	"github.com/maiaaraujo5/gostart/database/connection"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
)

type mongodb struct {
	client *mongo.Client
}

func NewMongoDB(config *database.Config) (database.Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), config.ConnectionTimeout)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.URI))
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}

	log.Printf("connect successfully in mongodb")
	return mongodb{
		client: client,
	}, nil
}

func (m mongodb) GetConnection(ctx context.Context) *connection.Connection {
	return &connection.Connection{
		MongoDB: m.client,
	}
}

func (m mongodb) Ping(ctx context.Context) error {
	err := m.client.Ping(ctx, readpref.Primary())
	if err != nil {
		return err
	}
	return nil
}

func (m mongodb) Disconnect(ctx context.Context) error {
	err := m.client.Disconnect(ctx)
	if err != nil {
		return err
	}
	return nil
}
