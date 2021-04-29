package connection

import "go.mongodb.org/mongo-driver/mongo"

type Connection struct {
	MongoDB *mongo.Client
}
