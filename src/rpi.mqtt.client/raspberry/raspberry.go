package raspberry

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type raspberry struct {
	client *mqtt.Client
	name   string
	debug  bool
}

func NewRaspberry(c *mqtt.Client, name string, debug bool) *raspberry {
	return &raspberry{c, name, debug}
}
