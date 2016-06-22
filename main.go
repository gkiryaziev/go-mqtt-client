package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	yamlCfg "github.com/gkiryaziev/go-config-manager/yaml"

	"github.com/gkiryaziev/go-mqtt-client/conf"
	"github.com/gkiryaziev/go-mqtt-client/raspberry"
	"github.com/gkiryaziev/go-mqtt-client/service"
)

// checkError check error
func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// config object
	var config conf.Config

	// config manager
	err := yamlCfg.NewConfig("config.yaml").Load(&config)
	checkError(err)

	// interrupt
	interrupt := make(chan os.Signal)
	signal.Notify(interrupt, os.Interrupt, os.Kill, syscall.SIGTERM)

	// open mqtt connection
	client, err := service.NewMqttClient(
		config.Mqtt.Protocol,
		config.Mqtt.Address,
		config.Mqtt.Port,
		0,
	)
	checkError(err)

	// new instance of mqtt client
	rpi := raspberry.NewRaspberry(client, config.Name, config.Debug)

	// Run publisher
	rpi.CPU.Themperature.Publish(config.Timeout, 0)
	rpi.CPU.CoreVolt.Publish(config.Timeout, 0)
	rpi.System.Memory.Publish(config.Timeout, 0)

	// Run subscribing
	rpi.System.Led.Subscribe(2)

	// wait for terminating
	for {
		select {
		case <-interrupt:
			log.Println("Clean and terminating...")

			// Unsubscribe when terminating
			rpi.System.Led.UnSubscribe()

			// disconnecting
			client.Disconnect(250)

			os.Exit(0)
		}
	}
}
