package database

import (
	"github.com/maiaaraujo5/gostart/config"
	"time"
)

type Config struct {
	Url               string
	Password          string
	DatabaseName      string
	URI               string
	ConnectionTimeout time.Duration
}

func NewConfig(conf config.Config) (*Config, error) {
	c := &Config{}

	if err := conf.ReadConfig(c, "db"); err != nil {
		return nil, err
	}

	return c, nil
}
