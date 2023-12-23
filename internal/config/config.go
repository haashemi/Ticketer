package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	APIAddr  string `yaml:"api_addr"` // RestAPI's listen address
	Database string `yaml:"database"` // Postgres' connection url
	JWTKey   string `yaml:"jwt_key"`  // JWT Tokens' key
	Host     string `yaml:"host"`     // Website's hostname
}

// Load returns the decoded Config from "./config.yaml" file.
func Load() (*Config, error) {
	data, err := os.ReadFile("./config.yaml")
	if err != nil {
		return nil, err
	}

	conf := &Config{}
	return conf, yaml.Unmarshal(data, conf)
}
