package conf

import "testing"

func TestLoad(t *testing.T) {
	// read and parse yaml file
	config, err := NewConfig("../../../config.yaml").Load()
	if err != nil {
		t.Fatal(err)
	}

	// check parameters
	switch {
	case config.Debug != true && config.Debug != false:
		t.Error("Error reading Debug parameter.")
	case config.Timeout == 0:
		t.Error("Error reading Timeout parameter.")
	case config.Name == "":
		t.Error("Error reading Name parameter.")
	case config.Mqtt.Protocol == "":
		t.Error("Error reading Protocol parameter.")
	case config.Mqtt.Address == "":
		t.Error("Error reading Address parameter.")
	case config.Mqtt.Port == "":
		t.Error("Error reading Port parameter.")
	}
}
