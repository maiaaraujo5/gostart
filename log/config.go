package log

import "github.com/maiaaraujo5/gostart/config"

const root = "gostart.log"

func DefaultConfig() {
	config.AddDefault(root+".impl", "zerolog")
}

func ImpValue() string {
	return config.GetStringValue(root + ".impl")
}
