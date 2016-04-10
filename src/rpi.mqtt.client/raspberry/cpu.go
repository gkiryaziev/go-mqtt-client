package raspberry

import (
	"log"
	"time"

	"rpi.mqtt.client/command"
	"rpi.mqtt.client/service/vcgencmd"
)

// Publish CPU temperature.
func (this *raspberry) CpuTemp(timeout int, qos byte) {

	topic := this.name+"/CPU/TEMP"

	log.Println("Start publishing: ", topic)

	time.Sleep(500 * time.Millisecond)

	// var temp string

	for {
		cpuTemp := vcgencmd.Clean(command.Exec("vcgencmd", "measure_temp"), "temp=", "'C")

		if cpuTemp != "" {
			// publish cpu temperature
			// if temp != cpuTemp {
				if token := this.client.Publish(topic, qos, false, cpuTemp); token.Wait() && token.Error() != nil {
					log.Println(token.Error())
				}

				// debug
				if this.debug {
					log.Println("[PUB]", topic, cpuTemp)
				}

				// temp = cpuTemp
			// }
		}

		time.Sleep(time.Duration(timeout) * time.Millisecond)
	}
}

// Publish CPU core volt.
func (this *raspberry) CpuCoreVolt(timeout int, qos byte) {

	topic := this.name+"/CPU/CORE/VOLT"

	log.Println("Start publishing: ", topic)

	time.Sleep(500 * time.Millisecond)

	// var volt string

	for {
		cpuCoreVolt := vcgencmd.Clean(command.Exec("vcgencmd", "measure_volts", "core"), "volt=", "V")

		if cpuCoreVolt != "" {
			// publish cpu core volt
			// if volt != cpuCoreVolt {
				if token := this.client.Publish(topic, qos, false, cpuCoreVolt); token.Wait() && token.Error() != nil {
					log.Println(token.Error())
				}

				// debug
				if this.debug {
					log.Println("[PUB]", topic, cpuCoreVolt)
				}

				// volt = cpuCoreVolt
			// }
		}

		time.Sleep(time.Duration(timeout) * time.Millisecond)
	}
}
