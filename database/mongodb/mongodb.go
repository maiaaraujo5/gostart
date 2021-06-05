package mongodb

import (
	"context"
	"github.com/maiaaraujo5/gostart/database"
	"github.com/maiaaraujo5/gostart/database/connection"
	"github.com/maiaaraujo5/gostart/log/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type mongodb struct {
	client *mongo.Client
}

func Connect() (database.Database, error) {
	config, err := NewConfig()
	if err != nil {
		return nil, err
	}

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
	logger.Info("connect successfully in mongodb")
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
