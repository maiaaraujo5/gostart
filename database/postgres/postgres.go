package postgres

import (
	"context"
	"github.com/go-pg/pg"
)

type postgres struct {
	client *pg.DB
}

func Connect() (*postgres, error) {

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

	return &postgres{
		client: client,
	}, nil
}

func (p *postgres) Disconnect(ctx context.Context) error {
	return p.client.Close()
}
