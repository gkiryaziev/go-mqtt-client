package syscore

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// System struct
type System struct {
	Memory *memory
	Led0   *led
	Led1   *led
}

// NewSystem return new System object.
func NewSystem(c mqtt.Client, name string, debug bool) *System {
	return &System{
		Memory: newMemory(c, name, debug),
		Led0:   newLed(c, name+"/SYSTEM/LED0", debug),
		Led1:   newLed(c, name+"/SYSTEM/LED1", debug),
	}
}
