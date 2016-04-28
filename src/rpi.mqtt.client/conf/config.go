package conf

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type config struct {
	file string
}

// NewConfig constructor
func NewConfig(file string) *config {
	return &config{file}
}

// Load config from file
func (this *config) Load() (*Config, error) {
	data, err := ioutil.ReadFile(this.file)
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

type Config struct {
	Debug   bool   `yaml:"debug"`
	Timeout int    `yaml:"timeout"`
	Name    string `yaml:"name"`
	Mqtt    Mqtt   `yaml:"mqtt"`
}

type Mqtt struct {
	Protocol string `yaml:"protocol"`
	Address  string `yaml:"address"`
	Port     string `yaml:"port"`
}
