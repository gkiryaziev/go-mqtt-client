package cpu_core

import (
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"rpi.mqtt.client/service/vcgencmd"
	"rpi.mqtt.client/command"
)

type CoreVolt struct {
	client *mqtt.Client
	debug  bool
	topic string
}

func newCoreVolt(c *mqtt.Client, name string, debug bool) *CoreVolt {
	return &CoreVolt {
		client: c,
		debug:  debug,
		topic:  name + "/CPU/CORE/VOLT",
	}
}

func (this *CoreVolt) Publish(timeout int, qos byte) {
	go func() {
		log.Println("Start publishing: ", this.topic)

		time.Sleep(500 * time.Millisecond)

		for {
			this.PublishOnce(qos)
			time.Sleep(time.Duration(timeout) * time.Millisecond)
		}
	}()
}

func (this *CoreVolt) PublishOnce(qos byte) {

	cpuCoreVolt := vcgencmd.Clean(command.Exec("vcgencmd", "measure_volts", "core"), "volt=", "V")

	if cpuCoreVolt != "" {

		if token := this.client.Publish(this.topic, qos, false, cpuCoreVolt); token.Wait() && token.Error() != nil {
			log.Println(token.Error())
		}

		// debug
		if this.debug {
			log.Println("[PUB]", this.topic, cpuCoreVolt)
		}
	}
}

func (this *CoreVolt) Subscribe() {}