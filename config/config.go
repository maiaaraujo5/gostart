package config

import (
	"github.com/spf13/viper"
	"os"
	"strings"
)

type Config interface {
	ReadConfig(i interface{}, key string) error
}

type config struct {
	Viper *viper.Viper
}

func NewConfig() (Config, error) {
	v := viper.New()
	conf := os.Getenv("CONF")
	files := strings.Split(conf, ",")

	for _, file := range files {
		v.SetConfigFile(file)

		err := v.MergeInConfig()
		if err != nil {
			return nil, err
		}
	}

	return &config{
		Viper: v,
	}, nil

}

func (v *config) ReadConfig(i interface{}, key string) error {
	err := v.Viper.UnmarshalKey(key, i)
	if err != nil {
		return err
	}

	return nil
}
