package rest

import "github.com/maiaaraujo5/gostart/config"

type Config struct {
	Port         string
	Cors         bool
	AllowOrigins []string
}

func NewConfig(conf config.Config) (*Config, error) {
	c := &Config{}

	if err := conf.ReadConfig(c, "rest"); err != nil {
		return nil, err
	}

	return c, nil
}
