package conf

import "testing"

func TestLoad(t *testing.T) {
	config, err := NewConfig("../../../config.yaml").Load()
	if err != nil {
		t.Fatal(err)
	}
	// check debug record
	if config.Debug != true && config.Debug != false {
		t.Error("Error reading config.")
	}
}
