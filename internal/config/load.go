package config

import (
	"gopkg.in/yaml.v3"
	"io"
	"os"
)

func LoadConfigFromFile(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	return LoadConfig(file)
}

func LoadConfig(reader io.Reader) (*Config, error) {
	decoder := yaml.NewDecoder(reader)
	config := &Config{}
	err := decoder.Decode(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
