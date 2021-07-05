package rabbitmq

import "github.com/maiaaraujo5/gostart/config"

const (
	root = "gostart.rabbitmq"
)

type Config struct {
	URL    string
	Sentry bool
}

func defaultConfig() {
	config.AddDefault(root+".url", "amqps://localhost:5672/temp")
	config.AddDefault(root+".sentry", false)
}

func NewConfig() (*Config, error) {
	defaultConfig()
	c := &Config{}

	err := config.ReadConfigPath(&c, root)
	if err != nil {
		return nil, err
	}

	return c, nil
}
