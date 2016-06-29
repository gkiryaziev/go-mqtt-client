package service

import (
	"bufio"
	"os"
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
