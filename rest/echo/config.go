package echo

import "github.com/maiaaraujo5/gostart/config"

type Config struct {
	Port string
}

func NewConfig(conf config.Config) (*Config, error) {
	c := &Config{}

	if err := conf.ReadConfig(c, "echo"); err != nil {
		return nil, err
	}

	return c, nil
}
