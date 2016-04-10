package sys_core

import (
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var (
	gDebug bool   // global debug
	gTopic string // global topic
)

type Led struct {
	client *mqtt.Client
	debug  bool
	topic  string
}

func newLed(c *mqtt.Client, name string, debug bool) *Led {

	// global variables for handler
	gDebug = debug
	gTopic = name + "/SYSTEM/LED0"

	return &Led{
		client: c,
		debug:  gDebug,
		topic:  gTopic,
	}
}

// Publish led status only once
func (this *Led) PublishOnce(qos byte) {

	// get led status
	status := "OFF"

	topic := this.topic + "/STATUS"

	// publish result
	if token := this.client.Publish(topic, qos, false, status); token.Wait() && token.Error() != nil {
		log.Println(token.Error())
	}

	// debug
	if this.debug {
		log.Println("[PUB]", topic, status)
	}
}

// Subscribe
func (this *Led) Subscribe(qos byte) {

	log.Println("[RUN] Subscribing:", this.topic)

	if token := this.client.Subscribe(this.topic + "/ACTION", qos, ledOnMessage); token.Wait() && token.Error() != nil {
		log.Println(token.Error())
	}
}

// LED0 onMessage handler
var ledOnMessage mqtt.MessageHandler = func(client *mqtt.Client, msg mqtt.Message) {

	// debug
	if gDebug {
		log.Println("[SUB]", msg.Topic(), string(msg.Payload()))
	}

	// receive message and DO
	status := "OFF"

	topic := gTopic + "/STATUS"

	// publish result
	if token := client.Publish(topic, msg.Qos(), false, status); token.Wait() && token.Error() != nil {
		log.Println(token.Error())
	}

	// debug
	if gDebug {
		log.Println("[PUB]", topic, status)
	}
}
