package config

import "github.com/maiaaraujo5/gostart/config"

const root = "gostart.echo"

type Config struct {
	Port         string
	Cors         bool
	AllowOrigins []string
	Sentry       bool
}

func defaultConfig() {
	config.AddDefault(root+".port", "8080")
	config.AddDefault(root+".cors", false)
	config.AddDefault(root+".alloworigins", []string{""})
	config.AddDefault(root+".sentry", false)
}

func NewConfig() (*Config, error) {
	c := &Config{}

	defaultConfig()

	if err := config.ReadConfigPath(c, root); err != nil {
		return nil, err
	}

	return c, nil
}

func SentryEnabled() bool {
	return config.GetBoolValue(root + ".sentry")
}
