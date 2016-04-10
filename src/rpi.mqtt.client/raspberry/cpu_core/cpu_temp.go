package cpu_core

import (
	"log"
	"time"

	"rpi.mqtt.client/command"
	"rpi.mqtt.client/service/vcgencmd"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Themperature struct {
	client *mqtt.Client
	debug  bool
	topic  string
}

func newThemperature(c *mqtt.Client, name string, debug bool) *Themperature {
	return &Themperature{
		client: c,
		debug:  debug,
		topic:  name + "/CPU/TEMP",
	}
}

func (this *Themperature) Publish(timeout int, qos byte) {
	go func() {
		log.Println("Start publishing: ", this.topic)

		time.Sleep(500 * time.Millisecond)

		for {
			this.PublishOnce(qos)
			time.Sleep(time.Duration(timeout) * time.Millisecond)
		}
	}()
}

func (this *Themperature) PublishOnce(qos byte) {

	cpuTemp := vcgencmd.Clean(command.Exec("vcgencmd", "measure_temp"), "temp=", "'C")

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

func (this *Themperature) Subscribe() {}
