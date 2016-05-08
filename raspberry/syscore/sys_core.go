package syscore

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// System struct
type System struct {
	Memory *memory
	Led    *led
}

// NewSystem return new System object.
func NewSystem(c mqtt.Client, name string, debug bool) *System {
	return &System{
		Memory: newMemory(c, name, debug),
		Led:    newLed(c, name, debug),
	}
}
