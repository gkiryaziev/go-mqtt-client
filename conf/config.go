package conf

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// ConfigManager struct
type ConfigManager struct {
	file string
}

// NewConfig constructor
func NewConfig(file string) *ConfigManager {
	return &ConfigManager{file}
}

// Load config from file
func (cm *ConfigManager) Load() (*Config, error) {
	data, err := ioutil.ReadFile(cm.file)
	if err != nil {
		return nil, err
	}

	var config *Config

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

// Config struct
type Config struct {
	Debug   bool   `yaml:"debug"`
	Timeout int    `yaml:"timeout"`
	Name    string `yaml:"name"`
	Mqtt    Mqtt   `yaml:"mqtt"`
}

// Mqtt struct
type Mqtt struct {
	Protocol string `yaml:"protocol"`
	Address  string `yaml:"address"`
	Port     string `yaml:"port"`
}
