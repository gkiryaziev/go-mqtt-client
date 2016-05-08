package raspberry

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"

	"github.com/gkiryaziev/go-mqtt-client/raspberry/cpucore"
	"github.com/gkiryaziev/go-mqtt-client/raspberry/syscore"
)

// Raspberry struct
type Raspberry struct {
	client mqtt.Client
	name   string
	CPU    *cpucore.CPU
	System *syscore.System
}

// NewRaspberry return new raspberry object.
func NewRaspberry(c mqtt.Client, name string, debug bool) *Raspberry {
	return &Raspberry{
		client: c,
		name:   name,
		CPU:    cpucore.NewCPU(c, name, debug),
		System: syscore.NewSystem(c, name, debug),
	}
}
