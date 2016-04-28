##	Golang MQTT Client for Raspberry PI 2

[![Go Report Card](https://goreportcard.com/badge/github.com/gkiryaziev/go_raspberry_pi_mqtt_client)](https://goreportcard.com/report/github.com/gkiryaziev/go_raspberry_pi_mqtt_client)
![Release Status](https://img.shields.io/badge/status-beta-yellow.svg?style=flat)


[Go](https://golang.org/) mqtt client with [Paho](http://www.eclipse.org/paho/clients/golang/) library.

## Installation

#### 1. Install GO
#### 2. Install GB
  `go get -u github.com/constabulary/gb/...`
#### 3. Clone project
  `git clone https://github.com/gkiryaziev/go_raspberry_pi_mqtt_client.git`
#### 4. Restore vendors
  `cd go_raspberry_pi_mqtt_client`
  
  `gb vendor restore`
#### 5. Edit configuration
  Copy `config.default.yaml` to `config.yaml` and edit configuration.
#### 6. Build and Run project
  `gb build && bin/rpi.mqtt.client run`