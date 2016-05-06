package raspberry

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"

	"github.com/gkiryaziev/go-mqtt-client/raspberry/cpu_core"
	"github.com/gkiryaziev/go-mqtt-client/raspberry/sys_core"
)

type raspberry struct {
	client mqtt.Client
	name   string
	CPU    *cpu_core.Cpu
	System *sys_core.System
}

// NewRaspberry return new raspberry object.
func NewRaspberry(c mqtt.Client, name string, debug bool) *raspberry {
	return &raspberry{
		client: c,
		name:   name,
		CPU:    cpu_core.NewCpu(c, name, debug),
		System: sys_core.NewSystem(c, name, debug),
	}
}
