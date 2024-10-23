package config

import (
	"os"

	"github.com/pelletier/go-toml/v2"
)

type Config struct {
	Database struct {
		User     string
		Password string
		Host     string
		Port     string
		Database string
	}

	HTTP struct {
		Port string
	}
}

func NewConfig(path string) (*Config, error) {
	config := new(Config)

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	err = toml.NewDecoder(file).Decode(config)
	if err != nil {
		panic(err)
	}

	return config, nil
}
