package raspberry

import (
	"log"
	"time"

	"rpi.mqtt.client/command"
	"rpi.mqtt.client/service/meminfo"
)

// Get system memory.
func (this *raspberry) SystemMemory(timeout int, qos byte) {

	topicMemTotal := this.name + "/SYSMEM/TOTAL"
	topicMemFree := this.name + "/SYSMEM/FREE"
	topicMemAvailable := this.name + "/SYSMEM/AVAILABLE"

	log.Println("Start publishing: ", topicMemTotal, topicMemFree, topicMemAvailable)

	time.Sleep(500 * time.Millisecond)

	var memTotal, memFree, memAvailable string

	for {
		sysMem := meminfo.Clean(command.Exec("cat", "/proc/meminfo"), "MemTotal:", "MemFree:", "MemAvailable:")

		if sysMem != nil {

			// publish total memory
			if memTotal != sysMem["MemTotal"] {
				if token := this.client.Publish(topicMemTotal, qos, false, sysMem["MemTotal"]); token.Wait() && token.Error() != nil {
					log.Println(token.Error())
				}

				// debug
				if this.debug {
					log.Println(topicMemTotal, sysMem["MemTotal"])
				}

				memTotal = sysMem["MemTotal"]
			}

			// publish free memory
			if memFree != sysMem["MemFree"] {
				if token := this.client.Publish(topicMemFree, qos, false, sysMem["MemFree"]); token.Wait() && token.Error() != nil {
					log.Println(token.Error())
				}

				// debug
				if this.debug {
					log.Println(topicMemFree, sysMem["MemFree"])
				}

				memFree = sysMem["MemFree"]
			}

			// publish available memory
			if memAvailable != sysMem["MemAvailable"] {
				if token := this.client.Publish(topicMemAvailable, qos, false, sysMem["MemAvailable"]); token.Wait() && token.Error() != nil {
					log.Println(token.Error())
				}

				// debug
				if this.debug {
					log.Println(topicMemAvailable, sysMem["MemAvailable"])
				}

				memAvailable = sysMem["MemAvailable"]
			}
		}

		time.Sleep(time.Duration(timeout) * time.Millisecond)
	}
}
