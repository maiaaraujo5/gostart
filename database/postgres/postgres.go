package postgres

import (
	"context"
	"github.com/go-pg/pg"
	"github.com/maiaaraujo5/gostart/database"
	"github.com/maiaaraujo5/gostart/database/connection"
)

type postgres struct {
	client *pg.DB
}

func Connect() (database.Database, error) {

	config, err := NewConfig()
	if err != nil {
		return nil, err
	}

	client := pg.Connect(&pg.Options{
		User:     config.User,
		Password: config.Password,
		Addr:     config.Addr,
		PoolSize: config.PoolSize,
	}).WithTimeout(config.ConnectionTimeout)

	return postgres{
		client: client,
	}, nil
}

func (p postgres) GetConnection(ctx context.Context) *connection.Connection {
	return &connection.Connection{
		Postgres: p.client,
	}
}

func (p postgres) Ping(ctx context.Context) error {
	panic("implement me")
}

func (p postgres) Disconnect(ctx context.Context) error {
	return p.client.Close()
}
