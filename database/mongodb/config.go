package mongodb

import (
	"github.com/maiaaraujo5/gostart/config"
	"time"
)

const root = "gostart.mongodb"

type Config struct {
	URI               string
	ConnectionTimeout time.Duration
}

func defaultConfig() {
	config.AddDefault(root+".uri", "mongodb://localhost:27017/temp")
	config.AddDefault(root+".connectiontimeout", 5*time.Second)
}

func NewConfig() (*Config, error) {
	c := &Config{}

	defaultConfig()

	if err := config.ReadConfigPath(c, root); err != nil {
		return nil, err
	}

	return c, nil
}
