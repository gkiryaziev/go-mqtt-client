package syscore

import (
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"

	"github.com/gkiryaziev/go-mqtt-client/service"
	"github.com/gkiryaziev/go-mqtt-client/service/meminfo"
)

type memory struct {
	client mqtt.Client
	debug  bool
	topic  string
}

// newMemory return new memory object.
func newMemory(c mqtt.Client, name string, debug bool) *memory {
	return &memory{
		client: c,
		debug:  debug,
		topic:  name + "/SYSTEM/MEMORY",
	}
}

// Publish system memory in goroutine with timeout
func (m *memory) Publish(timeout int, qos byte) {
	go func() {
		log.Println("[RUN] Publishing:", qos, m.topic)

		time.Sleep(500 * time.Millisecond)

		for {
			m.PublishOnce(qos)
			time.Sleep(time.Duration(timeout) * time.Millisecond)
		}
	}()
}

// PublishOnce publish system memory only once
func (m *memory) PublishOnce(qos byte) {

	topicMemTotal := m.topic + "/TOTAL"
	topicMemFree := m.topic + "/FREE"
	topicMemAvailable := m.topic + "/AVAILABLE"

	sysMem := meminfo.Clean(service.CmdExec("cat", "/proc/meminfo"), "MemTotal:", "MemFree:", "MemAvailable:")

	if sysMem != nil {

		// publish total memory
		if token := m.client.Publish(topicMemTotal, qos, false, sysMem["MemTotal"]); token.Wait() && token.Error() != nil {
			log.Println(token.Error())
		}

		// debug
		if m.debug {
			log.Println("[PUB]", qos, topicMemTotal, sysMem["MemTotal"])
		}

		// publish free memory
		if token := m.client.Publish(topicMemFree, qos, false, sysMem["MemFree"]); token.Wait() && token.Error() != nil {
			log.Println(token.Error())
		}

		// debug
		if m.debug {
			log.Println("[PUB]", qos, topicMemFree, sysMem["MemFree"])
		}

		// publish available memory
		if token := m.client.Publish(topicMemAvailable, qos, false, sysMem["MemAvailable"]); token.Wait() && token.Error() != nil {
			log.Println(token.Error())
		}

		// debug
		if m.debug {
			log.Println("[PUB]", qos, topicMemAvailable, sysMem["MemAvailable"])
		}
	}
}

// Subscribe to topic
func (m *memory) Subscribe() {}
