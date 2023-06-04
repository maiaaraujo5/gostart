package sentry

import (
	"github.com/maiaaraujo5/gostart/config"
	"github.com/maiaaraujo5/gostart/echo/plugins"
	"time"
)

const (
	root            = plugins.Root + ".sentry"
	name            = root + ".name"
	description     = root + ".description"
	enabled         = root + ".enabled"
	dsn             = root + ".dsn"
	environment     = root + ".environment"
	traceSampleRate = root + ".tracesamplerate"
	timeout         = root + ".timeout"
	debug           = root + ".debug"
)

type Config struct {
	Name            string
	Description     string
	Debug           bool
	Environment     string
	Enabled         bool
	Dsn             string
	TraceSampleRate float64
	Timeout         time.Duration
}

func defaultConfig() {
	config.AddDefault(enabled, false)
	config.AddDefault(name, "sentry")
	config.AddDefault(description, "sentry health")
	config.AddDefault(environment, "homolog")
	config.AddDefault(dsn, "")
	config.AddDefault(traceSampleRate, 1.0)
	config.AddDefault(timeout, 3*time.Second)
	config.AddDefault(debug, false)
}

func NewConfig() (*Config, error) {
	c := &Config{}

	defaultConfig()

	err := config.ReadConfigPath(c, root)
	if err != nil {
		return nil, err
	}

	return c, nil
}
