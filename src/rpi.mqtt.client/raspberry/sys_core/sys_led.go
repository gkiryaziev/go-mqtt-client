package sys_core

import (
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Led struct {
	client mqtt.Client
	debug  bool
	topic  string
}

func newLed(c mqtt.Client, name string, debug bool) *Led {
	return &Led{
		client: c,
		debug:  debug,
		topic:  name + "/SYSTEM/LED0",
	}
}

// Subscribe
func (this *Led) Subscribe(qos byte) {

	topic := this.topic + "/ACTION"

	log.Println("[RUN] Subscribing:", qos, topic)

	if token := this.client.Subscribe(topic, qos, this.ledOnMessage); token.Wait() && token.Error() != nil {
		log.Println(token.Error())
	}
}

// Subscribe
func (this *Led) UnSubscribe() {

	topic := this.topic + "/ACTION"

	log.Println("[RUN] UnSubscribing:", topic)

	if token := this.client.Unsubscribe(topic); token.Wait() && token.Error() != nil {
		log.Println(token.Error())
	}
}

// LED0 onMessage handler
func (this *Led) ledOnMessage(client mqtt.Client, msg mqtt.Message) {

	// debug
	if this.debug {
		log.Println("[SUB]", msg.Qos(), msg.Topic(), string(msg.Payload()))
	}

	// receive message and DO
	status := "OFF"

	switch string(msg.Payload()) {
	case "ON":
		// logic when ON
		status = "ON"
	case "OFF":
		// logic when OFF
		status = "OFF"
	}

	topic := this.topic + "/STATUS"

	// publish result
	if token := client.Publish(topic, msg.Qos(), false, status); token.Wait() && token.Error() != nil {
		log.Println(token.Error())
	}

	// debug
	if this.debug {
		log.Println("[PUB]", msg.Qos(), topic, status)
	}
}
