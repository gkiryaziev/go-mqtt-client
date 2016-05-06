package cpu_core

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Cpu struct {
	Themperature *themperature
	CoreVolt     *coreVolt
}

// NewCpu return new Cpu object.
func NewCpu(c mqtt.Client, name string, debug bool) *Cpu {
	return &Cpu{
		Themperature: newThemperature(c, name, debug),
		CoreVolt:     newCoreVolt(c, name, debug),
	}
}
