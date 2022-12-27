package echo

import "github.com/maiaaraujo5/gostart/config"

const Root = "gostart.echo"

type Config struct {
	Port         string
	Cors         bool
	AllowOrigins []string
}

func defaultConfig() {
	config.AddDefault(Root+".port", "8080")
	config.AddDefault(Root+".cors", true)
	config.AddDefault(Root+".alloworigins", []string{""})
}

func NewConfig() (*Config, error) {
	c := &Config{}

	defaultConfig()

	if err := config.ReadConfigPath(c, Root); err != nil {
		return nil, err
	}

	return c, nil
}
