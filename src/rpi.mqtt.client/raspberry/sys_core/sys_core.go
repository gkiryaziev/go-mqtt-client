package sys_core

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type System struct {
	Memory *Memory
}

func NewSystem(c *mqtt.Client, name string, debug bool) *System {
	return &System{
		Memory: newMemory(c, name, debug),
	}
}
