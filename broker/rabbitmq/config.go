package rabbitmq

import "github.com/maiaaraujo5/gostart/config"

const (
	root = "gostart.rabbitmq"
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

func defaultConfig() {
	config.AddDefault(root+".url", "amqps://localhost:5672/temp")
}

func NewConfig() (*Config, error) {
	c := &Config{}

	err := config.ReadConfigPath(&c, root)
	if err != nil {
		return nil, err
	}

	return c, nil
}
