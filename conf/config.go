// Package conf for go-config-manager
package conf

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
