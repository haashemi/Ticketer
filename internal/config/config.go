package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	APIAddr  string `yaml:"api_addr"`
	Database string `yaml:"database"`
}

func Load() (*Config, error) {
	data, err := os.ReadFile("./config.yaml")
	if err != nil {
		return nil, err
	}

	conf := &Config{}
	return conf, yaml.Unmarshal(data, conf)
}
