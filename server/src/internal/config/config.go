package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type ServiceConfig struct {
	Server struct {
		Port uint16 `yaml:"port"`
	} `yaml:"server"`

	InfluenceDB PgConfig `yaml:"influencedb"`

	Common struct {
		PlayerNameRegex string `yaml:"player_name_regex"`
	} `yaml:"common"`
}

func FromYaml(path string) (ServiceConfig, error) {
	content, err := os.ReadFile(path)

	if err != nil {
		return ServiceConfig{}, fmt.Errorf("Cannot read config file: %e", err)
	}

	config := ServiceConfig{}
	err = yaml.Unmarshal(content, &config)

	if err != nil {
		return ServiceConfig{}, fmt.Errorf("Cannot parse config file: %e", err)
	}

	return config, nil
}
