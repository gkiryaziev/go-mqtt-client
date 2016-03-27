package command

import (
	"log"
	"os/exec"
)

// exec command
func Exec(name string, args ...string) string {
	out, err := exec.Command(name, args...).Output()
	if err != nil {
		log.Println(err)
		return ""
	}
	return string(out)
}
