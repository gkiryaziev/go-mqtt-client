package syscore

import (
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type led struct {
	client mqtt.Client
	debug  bool
	topic  string
}

// NewSystem return new led object.
func newLed(c mqtt.Client, name string, debug bool) *led {
	return &led{
		client: c,
		debug:  debug,
		topic:  name + "/SYSTEM/LED0",
	}
}

// Subscribe to topic
func (l *led) Subscribe(qos byte) {

	topic := l.topic + "/ACTION"

	log.Println("[RUN] Subscribing:", qos, topic)

	if token := l.client.Subscribe(topic, qos, l.ledOnMessage); token.Wait() && token.Error() != nil {
		log.Println(token.Error())
	}
}

// UnSubscribe from topic
func (l *led) UnSubscribe() {

	topic := l.topic + "/ACTION"

	log.Println("[RUN] UnSubscribing:", topic)

	if token := l.client.Unsubscribe(topic); token.Wait() && token.Error() != nil {
		log.Println(token.Error())
	}
}

// ledOnMessage set Led to on or of
func (l *led) ledOnMessage(client mqtt.Client, msg mqtt.Message) {

	// debug
	if l.debug {
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

	topic := l.topic + "/STATUS"

	// publish result
	if token := client.Publish(topic, msg.Qos(), false, status); token.Wait() && token.Error() != nil {
		log.Println(token.Error())
	}

	// debug
	if l.debug {
		log.Println("[PUB]", msg.Qos(), topic, status)
	}
}
