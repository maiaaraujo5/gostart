package health

import (
	"context"
	"github.com/maiaaraujo5/gostart/database/mongodb"
	"github.com/maiaaraujo5/gostart/health"
	"github.com/maiaaraujo5/gostart/log/logger"
	"go.mongodb.org/mongo-driver/mongo"
)

type Health struct {
	client *mongo.Client
}

func NewHealth(ctx context.Context) mongodb.Plugin {
	logger.Trace("registering health mongo plugin")

	c, err := NewConfig()
	if err != nil {
		logger.Fatal("error to get config from health mongo plugin")
	}

	return func(ctx context.Context, client *mongo.Client) error {
		h := &Health{client: client}
		check := health.NewHealth(c.Name, c.Description, h, c.Enabled)
		health.Add(check)

		logger.Debug("successfully registered health mongo plugin")
		return nil
	}
}

func (h *Health) Check(ctx context.Context) error {
	return h.client.Ping(ctx, nil)
}
