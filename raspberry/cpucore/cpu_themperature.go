package cpucore

import (
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"

	"github.com/gkiryaziev/go-mqtt-client/service"
	"github.com/gkiryaziev/go-mqtt-client/service/vcgencmd"
)

type themperature struct {
	client mqtt.Client
	debug  bool
	topic  string
}

// newThemperature return new themperature object.
func newThemperature(c mqtt.Client, name string, debug bool) *themperature {
	return &themperature{
		client: c,
		debug:  debug,
		topic:  name + "/CPU/TEMP",
	}
}

// Publish core themperature in goroutine with timeout
func (t *themperature) Publish(timeout int, qos byte) {
	go func() {
		log.Println("[RUN] Publishing:", qos, t.topic)

		time.Sleep(500 * time.Millisecond)

		for {
			t.PublishOnce(qos)
			time.Sleep(time.Duration(timeout) * time.Millisecond)
		}
	}()
}

// PublishOnce core themperature only once
func (t *themperature) PublishOnce(qos byte) {

	cpuTemp := vcgencmd.Clean(service.CmdExec("vcgencmd", "measure_temp"), "temp=", "'C")

	if cpuTemp != "" {

		if token := t.client.Publish(t.topic, qos, false, cpuTemp); token.Wait() && token.Error() != nil {
			log.Println(token.Error())
		}

		// debug
		if t.debug {
			log.Println("[PUB]", qos, t.topic, cpuTemp)
		}

	}
}

// Subscribe to topic
func (t *themperature) Subscribe() {}
