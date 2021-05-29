package rest

import "github.com/maiaaraujo5/gostart/config"

const root = "gostart.rest"

type Config struct {
	Port         string
	Cors         bool
	AllowOrigins []string
}

func defaultConfig() {
	config.AddDefault(root+".port", "8080")
	config.AddDefault(root+".cors", false)
	config.AddDefault(root+".alloworigins", []string{""})
}

func NewConfig() (*Config, error) {
	c := &Config{}

	defaultConfig()

	if err := config.ReadConfigPath(c, root); err != nil {
		return nil, err
	}

	return c, nil
}
