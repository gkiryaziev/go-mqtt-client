package raspberry

import (
	"rpi.mqtt.client/raspberry/cpu_core"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type raspberry struct {
	client *mqtt.Client
	name   string
	debug  bool
	Cpu    *cpu_core.Cpu
}

func NewRaspberry(c *mqtt.Client, name string, debug bool) *raspberry {
	return &raspberry{
		client: c,
		name:   name,
		debug:  debug,
		Cpu:    cpu_core.NewCpu(c, name, debug),
	}
}
