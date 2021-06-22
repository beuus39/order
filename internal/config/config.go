package config

import (
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	Database struct{
		DBName string   `yaml:"db_name"`
		DBUser string   `yaml:"db_user"`
		Host string     `yaml:"host"`
		Password string `yaml:"password"`
		Dialect string  `yaml:"dialect"`
	}
}

func NewConfig(configFile string) (*Config, error) {
	file, err := os.Open(configFile)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	cfg := &Config{}

	yd := yaml.NewDecoder(file)
	err = yd.Decode(cfg)

	if err != nil {
		return nil, err
	}
	return cfg, nil
}
