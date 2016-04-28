package cpu_core

import (
	"log"
	"time"

	"rpi.mqtt.client/service"
	"rpi.mqtt.client/service/vcgencmd"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type coreVolt struct {
	client mqtt.Client
	debug  bool
	topic  string
}

func newCoreVolt(c mqtt.Client, name string, debug bool) *coreVolt {
	return &coreVolt{
		client: c,
		debug:  debug,
		topic:  name + "/CPU/CORE/VOLT",
	}
}

// Publish core volt in goroutine with timeout
func (this *coreVolt) Publish(timeout int, qos byte) {
	go func() {
		log.Println("[RUN] Publishing:", qos, this.topic)

		time.Sleep(500 * time.Millisecond)

		for {
			this.PublishOnce(qos)
			time.Sleep(time.Duration(timeout) * time.Millisecond)
		}
	}()
}

// PublishOnce core volt only once
func (this *coreVolt) PublishOnce(qos byte) {

	cpuCoreVolt := vcgencmd.Clean(service.CmdExec("vcgencmd", "measure_volts", "core"), "volt=", "V")

	if cpuCoreVolt != "" {

		if token := this.client.Publish(this.topic, qos, false, cpuCoreVolt); token.Wait() && token.Error() != nil {
			log.Println(token.Error())
		}

		// debug
		if this.debug {
			log.Println("[PUB]", qos, this.topic, cpuCoreVolt)
		}
	}
}

// Subscribe to topic
func (this *coreVolt) Subscribe() {}
