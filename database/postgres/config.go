package postgres

import (
	"github.com/maiaaraujo5/gostart/config"
	"time"
)

const root = "gostart.postgres"

type Config struct {
	User              string
	Password          string
	Addr              string
	PoolSize          int
	ConnectionTimeout time.Duration
}

func defaultConfig() {
	config.AddDefault(root+".user", "admin")
	config.AddDefault(root+".password", "admin")
	config.AddDefault(root+".addr", "localhost:5432")
	config.AddDefault(root+".poolsize", 1)
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
