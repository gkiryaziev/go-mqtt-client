package conf

import "testing"

func TestLoad(t *testing.T) {
	// read and parse yaml file
	config, err := NewConfig("../config.yaml").Load()
	if err != nil {
		t.Fatal(err)
	}

	// check parameters
	switch {
	case config.Debug != true && config.Debug != false:
		t.Error("Error, Debug =", config.Debug)
	case config.Timeout == 0:
		t.Error("Error, Timeout = 0")
	case config.Name == "":
		t.Error("Error, Name is empty.")
	case config.Mqtt.Protocol != "tcp":
		t.Error("Error, Protocol =", config.Mqtt.Protocol)
	case config.Mqtt.Address == "":
		t.Error("Error, Address is empty.")
	case config.Mqtt.Port == "":
		t.Error("Error, Port is empty.")
	}
}
