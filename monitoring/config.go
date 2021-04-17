package monitoring

import "github.com/maiaaraujo5/gostart/config"

type Config struct {
	Release           string
	Environment       string
	AuthenticationKey string
}

func NewConfig(conf config.Config) (*Config, error) {
	c := &Config{}

	if err := conf.ReadConfig(c, "monitoring"); err != nil {
		return nil, err
	}

	return c, nil
}
