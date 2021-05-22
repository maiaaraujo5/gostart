package rabbitmq

import "github.com/maiaaraujo5/gostart/config"

const (
	brokerConfig = "gostart.rabbitmq"
)

type Config struct {
	URL      string
	Sender   []configs
	Listener configs
}

type configs struct {
	Exchange   string
	Queue      string
	RoutingKey string
	Mandatory  bool
	Immediate  bool
}

func NewConfig(config config.Config) (*Config, error) {
	c := &Config{}

	err := config.ReadConfig(&c, brokerConfig)
	if err != nil {
		return nil, err
	}

	return c, nil
}
