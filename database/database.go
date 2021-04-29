package database

import (
	"context"
	"github.com/maiaaraujo5/gostart/database/connection"
)

type Database interface {
	GetConnection(ctx context.Context) *connection.Connection
	Ping(ctx context.Context) error
	Disconnect(ctx context.Context) error
}
