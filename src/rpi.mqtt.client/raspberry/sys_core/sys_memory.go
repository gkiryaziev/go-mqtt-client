package sys_core

import (
	"log"
	"time"

	"rpi.mqtt.client/service"
	"rpi.mqtt.client/service/meminfo"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Memory struct {
	client *mqtt.Client
	debug  bool
	topic  string
}

func newMemory(c *mqtt.Client, name string, debug bool) *Memory {
	return &Memory{
		client: c,
		debug:  debug,
		topic:  name + "/SYSMEM",
	}
}

// Publish system memory in goroutine with timeout
func (this *Memory) Publish(timeout int, qos byte) {
	go func() {
		log.Println("[RUN] Publishing:", this.topic)

		time.Sleep(500 * time.Millisecond)

		for {
			this.PublishOnce(qos)
			time.Sleep(time.Duration(timeout) * time.Millisecond)
		}
	}()
}

// Publish system memory only once
func (this *Memory) PublishOnce(qos byte) {

	topicMemTotal := this.topic + "/TOTAL"
	topicMemFree := this.topic + "/FREE"
	topicMemAvailable := this.topic + "/AVAILABLE"

	sysMem := meminfo.Clean(service.CmdExec("cat", "/proc/meminfo"), "MemTotal:", "MemFree:", "MemAvailable:")

	if sysMem != nil {

		// publish total memory
		if token := this.client.Publish(topicMemTotal, qos, false, sysMem["MemTotal"]); token.Wait() && token.Error() != nil {
			log.Println(token.Error())
		}

		// debug
		if this.debug {
			log.Println("[PUB]", topicMemTotal, sysMem["MemTotal"])
		}

		// publish free memory
		if token := this.client.Publish(topicMemFree, qos, false, sysMem["MemFree"]); token.Wait() && token.Error() != nil {
			log.Println(token.Error())
		}

		// debug
		if this.debug {
			log.Println("[PUB]", topicMemFree, sysMem["MemFree"])
		}

		// publish available memory
		if token := this.client.Publish(topicMemAvailable, qos, false, sysMem["MemAvailable"]); token.Wait() && token.Error() != nil {
			log.Println(token.Error())
		}

		// debug
		if this.debug {
			log.Println("[PUB]", topicMemAvailable, sysMem["MemAvailable"])
		}
	}
}

// Subscribe
func (this *Memory) Subscribe() {}
