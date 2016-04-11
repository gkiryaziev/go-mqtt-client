package cpu_core

import (
	"log"
	"time"

	"rpi.mqtt.client/service"
	"rpi.mqtt.client/service/vcgencmd"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Themperature struct {
	client mqtt.Client
	debug  bool
	topic  string
}

func newThemperature(c mqtt.Client, name string, debug bool) *Themperature {
	return &Themperature{
		client: c,
		debug:  debug,
		topic:  name + "/CPU/TEMP",
	}
}

// Publish core themperature in goroutine with timeout
func (this *Themperature) Publish(timeout int, qos byte) {
	go func() {
		log.Println("[RUN] Publishing:", this.topic)

		time.Sleep(500 * time.Millisecond)

		for {
			this.PublishOnce(qos)
			time.Sleep(time.Duration(timeout) * time.Millisecond)
		}
	}()
}

// Publish core themperature only once
func (this *Themperature) PublishOnce(qos byte) {

	cpuTemp := vcgencmd.Clean(service.CmdExec("vcgencmd", "measure_temp"), "temp=", "'C")

	if cpuTemp != "" {

		if token := this.client.Publish(this.topic, qos, false, cpuTemp); token.Wait() && token.Error() != nil {
			log.Println(token.Error())
		}

		// debug
		if this.debug {
			log.Println("[PUB]", this.topic, cpuTemp)
		}

	}
}

// Subscribe
func (this *Themperature) Subscribe() {}
