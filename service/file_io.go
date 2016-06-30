// Package service - temporary data for future usage
package service

import (
	"bufio"
	"io/ioutil"
	"os"
	"strings"
)

// ReadFile a size of byte from file
func ReadFile(file string, size int) (string, error) {
	f, err := os.Open(file)
	if err != nil {
		return "", err
	}
	defer f.Close()

	r := bufio.NewReader(f)
	b, err := r.Peek(size)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

// WriteFile some data to file
func WriteFile(file, data string) error {
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	_, err = w.WriteString(data)
	if err != nil {
		return err
	}

	w.Flush()

	return nil
}

// ReadTrigger read led trigger
func ReadTrigger(id string) (string, error) {
	dat, err := ioutil.ReadFile("/sys/class/leds/led" + id + "/trigger")
	if err != nil {
		return "", err
	}

	i1 := strings.Index(string(dat), "[")
	i2 := strings.Index(string(dat), "]")

	return string(dat)[i1+1 : i2], nil
}

// ReadBrightness read led brightness
func ReadBrightness(id string) (string, error) {
	dat, err := ioutil.ReadFile("/sys/class/leds/led" + id + "/brightness")
	if err != nil {
		return "", err
	}
	return string(dat), nil
}

// WriteBrightness write led brightness
func WriteBrightness(id, data string) error {
	err := ioutil.WriteFile("/sys/class/leds/led"+id+"/brightness", []byte(data), 0644)
	if err != nil {
		return err
	}
	return nil
}
