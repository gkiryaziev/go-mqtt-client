package service

import (
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/pborman/uuid"
)

// open mqtt connection
func NewMqttClient(protocol, address, port string, qos byte) (*mqtt.Client, error) {
	// generate new uuid
	id := uuid.New()
	// concat address
	server := protocol + "://" + address + ":" + port

	clientOptions := mqtt.NewClientOptions()
	clientOptions.AddBroker(server)
	clientOptions.SetClientID(id)
	clientOptions.SetDefaultPublishHandler(defaultMessageHandler)

	// new client
	client := mqtt.NewClient(clientOptions)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}

	return client, nil
}

var defaultMessageHandler mqtt.MessageHandler = func(client *mqtt.Client, msg mqtt.Message) {
	log.Printf("TOPIC: %s, MSG: %s\n", msg.Topic(), msg.Payload())
}
