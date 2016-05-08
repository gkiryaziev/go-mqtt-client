package cpucore

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// CPU struct
type CPU struct {
	Themperature *themperature
	CoreVolt     *coreVolt
}

// NewCPU return new Cpu object.
func NewCPU(c mqtt.Client, name string, debug bool) *CPU {
	return &CPU{
		Themperature: newThemperature(c, name, debug),
		CoreVolt:     newCoreVolt(c, name, debug),
	}
}
