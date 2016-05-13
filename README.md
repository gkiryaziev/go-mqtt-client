##	Golang MQTT Client for Raspberry PI 2

[![Go Report Card](https://goreportcard.com/badge/github.com/gkiryaziev/go-mqtt-client)](https://goreportcard.com/report/github.com/gkiryaziev/go-mqtt-client)

[Go](https://golang.org/) mqtt client with [Paho M2M](http://www.eclipse.org/paho/clients/golang/) library.

### Installation:
```
go get github.com/gkiryaziev/go-mqtt-client
```

### Edit configuration:
```
Copy `config.default.yaml` to `config.yaml` and edit configuration.
```

### Build and Run:
```
go build && go-mqtt-client
```

### Packages:
You can use [glide](https://glide.sh/) packages manager to get all needed packages.
```
go get -u -v github.com/Masterminds/glide

cd go-mqtt-client && glide install
```