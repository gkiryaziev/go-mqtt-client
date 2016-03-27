package vcgencmd

import "strings"

// vcgencmd result cleaner
func Clean(str string, args ...string) string {
	for _, arg := range args {
		str = strings.Replace(str, arg, "", -1)
	}
	str = strings.TrimSpace(str)
	return str
}

