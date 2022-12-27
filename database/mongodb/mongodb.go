package mongodb

import (
	"context"

	"github.com/maiaaraujo5/gostart/log/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Plugin func(ctx context.Context, client *mongo.Client) error

type mongodb struct {
	Client  *mongo.Client
	Plugins []Plugin
}

func Connect(ctx context.Context, plugins ...Plugin) (*mongodb, error) {

	config, err := NewConfig()
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(ctx, config.ConnectionTimeout)
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

	for _, pluginFunc := range plugins {
		err := pluginFunc(ctx, client)
		if err != nil {
			logger.Fatal(err.Error())
		}
	}

	return &mongodb{
		Client:  client,
		Plugins: plugins,
	}, nil
}

func (m *mongodb) Disconnect(ctx context.Context) error {
	err := m.Client.Disconnect(ctx)
	if err != nil {
		return err
	}
	return nil
}
