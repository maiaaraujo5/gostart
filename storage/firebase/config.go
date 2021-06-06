package firebase

import (
	"github.com/maiaaraujo5/gostart/config"
	"time"
)

const root = "gostart.firebase.storage"

type Config struct {
	StorageBucket         string
	PathToCredentialsFile string
	UploadTimeout         time.Duration
}

func defaultConfig() {
	config.AddDefault(root+".storagebucket", "tempbucket")
	config.AddDefault(root+".pathtocredentials", "./serviceAccountKey.json")
	config.AddDefault(root+".uploadtimeout", 30*time.Second)
}

func NewConfig() (*Config, error) {
	c := &Config{}

	if err := config.ReadConfigPath(&c, root); err != nil {
		return nil, err
	}

	return c, nil
}
