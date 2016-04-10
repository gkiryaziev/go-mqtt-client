package raspberry

import (
	"rpi.mqtt.client/raspberry/cpu_core"
	"rpi.mqtt.client/raspberry/sys_core"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type raspberry struct {
	client *mqtt.Client
	name   string
	debug  bool
	Cpu    *cpu_core.Cpu
	System *sys_core.System
}

func NewRaspberry(c *mqtt.Client, name string, debug bool) *raspberry {
	return &raspberry{
		client: c,
		name:   name,
		debug:  debug,
		Cpu:    cpu_core.NewCpu(c, name, debug),
		System: sys_core.NewSystem(c, name, debug),
	}
}
