package connection

import (
	"github.com/go-pg/pg"
	"go.mongodb.org/mongo-driver/mongo"
)

type Connection struct {
	MongoDB  *mongo.Client
	Postgres *pg.DB
}
