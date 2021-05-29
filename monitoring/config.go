package monitoring

import "github.com/maiaaraujo5/gostart/config"

const root = "gostart.monitoring"

type Config struct {
	Release           string
	Environment       string
	AuthenticationKey string
}

func NewConfig() (*Config, error) {
	c := &Config{}

	if err := config.ReadConfigPath(c, root); err != nil {
		return nil, err
	}

	return c, nil
}
