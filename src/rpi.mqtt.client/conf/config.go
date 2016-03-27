package conf

import (
	"github.com/jinzhu/configor"
)

var (
	config *Config
)

func init() {
	configor.Load(&config, "config.yaml")
}

func GetConfig() *Config {
	return config
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
