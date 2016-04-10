package raspberry

import (
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var (
	pTopic, sTopic string	// publish and subscribe topic
	gQos byte = 0			// global QoS
	gDebug bool = false		// global debug
)

func (this *raspberry) Led0(qos byte) error {

	//// unsubscribe
	//defer func() {
	//	if token := this.client.Unsubscribe(sTopic); token.Wait() && token.Error() != nil {
	//		log.Println(token.Error())
	//	}
	//}()

	sTopic = this.name + "/LED0/ACTION"
	pTopic = this.name + "/LED0/STATUS"
	gQos   = qos
	gDebug = this.debug

	log.Println("Start subscribing: ", sTopic)

	if token := this.client.Subscribe(sTopic, gQos, led0Handler); token.Wait() && token.Error() != nil {
		return token.Error()
	}

	return nil
}

var led0Handler mqtt.MessageHandler = func(client *mqtt.Client, msg mqtt.Message) {

	// receive message and DO

	// debug
	if gDebug {
		log.Println("[SUB]", msg.Topic(), string(msg.Payload()))
	}

	// publish result
	if token := client.Publish(pTopic, gQos, false, "OFF"); token.Wait() && token.Error() != nil {
		log.Println(token.Error())
	}

	// debug
	if gDebug {
		log.Println("[PUB]", pTopic, "OFF")
	}
}