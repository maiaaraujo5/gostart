package health

import (
	"github.com/maiaaraujo5/gostart/config"
	"github.com/maiaaraujo5/gostart/database/mongodb/plugins"
)

const (
	root        = plugins.Root + ".health"
	name        = root + ".name"
	description = root + ".description"
	enabled     = root + ".enabled"
)

type Config struct {
	Name        string
	Description string
	Enabled     bool
}

func defaultConfig() {
	config.AddDefault(name, "mongo")
	config.AddDefault(description, "mongo health")
	config.AddDefault(enabled, true)
}

func NewConfig() (*Config, error) {
	c := &Config{}

	defaultConfig()

	if err := config.ReadConfigPath(c, root); err != nil {
		return nil, err
	}

	return c, nil
}
