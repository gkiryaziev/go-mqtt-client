package service

import (
	"log"
	"os/exec"
)

// CmdExec execute command and return stdout
func CmdExec(name string, args ...string) string {
	out, err := exec.Command(name, args...).Output()
	if err != nil {
		log.Println(err)
		return ""
	}
	return string(out)
}
