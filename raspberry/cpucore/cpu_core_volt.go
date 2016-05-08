package cpucore

import (
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"

	"github.com/gkiryaziev/go-mqtt-client/service"
	"github.com/gkiryaziev/go-mqtt-client/service/vcgencmd"
)

type coreVolt struct {
	client mqtt.Client
	debug  bool
	topic  string
}

// newCoreVolt return new coreVolt object.
func newCoreVolt(c mqtt.Client, name string, debug bool) *coreVolt {
	return &coreVolt{
		client: c,
		debug:  debug,
		topic:  name + "/CPU/CORE/VOLT",
	}
}

// Publish core volt in goroutine with timeout
func (cv *coreVolt) Publish(timeout int, qos byte) {
	go func() {
		log.Println("[RUN] Publishing:", qos, cv.topic)

		time.Sleep(500 * time.Millisecond)

		for {
			cv.PublishOnce(qos)
			time.Sleep(time.Duration(timeout) * time.Millisecond)
		}
	}()
}

// PublishOnce core volt only once
func (cv *coreVolt) PublishOnce(qos byte) {

	cpuCoreVolt := vcgencmd.Clean(service.CmdExec("vcgencmd", "measure_volts", "core"), "volt=", "V")

	if cpuCoreVolt != "" {

		if token := cv.client.Publish(cv.topic, qos, false, cpuCoreVolt); token.Wait() && token.Error() != nil {
			log.Println(token.Error())
		}

		// debug
		if cv.debug {
			log.Println("[PUB]", qos, cv.topic, cpuCoreVolt)
		}
	}
}

// Subscribe to topic
func (cv *coreVolt) Subscribe() {}
