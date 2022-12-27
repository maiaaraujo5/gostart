package health

import (
	"github.com/maiaaraujo5/gostart/config"
	"github.com/maiaaraujo5/gostart/echo/plugins"
)

const (
	root    = plugins.Root
	enabled = root + ".enabled"
	route   = root + ".route"
)

type Config struct {
	Enabled bool
	Route   string
}

func defaultConfig() {
	config.AddDefault(enabled, true)
	config.AddDefault(route, "/health")
}

func NewConfig() (*Config, error) {
	c := &Config{}

	defaultConfig()

	if err := config.ReadConfigPath(c, root); err != nil {
		return nil, err
	}

	return c, nil
}
