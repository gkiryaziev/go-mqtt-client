package meminfo

import "strings"

// Clean cat /proc/meminfo cleaner (MemTotal: MemFree: MemAvailable:)
func Clean(str string, args ...string) map[string]string {
	if str == "" {
		return nil
	}
	result := make(map[string]string)
	strArray := strings.Split(str, "\n")
	for _, arg := range args {
		for _, val := range strArray {
			if str := strings.Split(val, arg); len(str) == 2 {
				newStr := strings.TrimSpace(strings.Replace(str[1], "kB", "", -1))
				result[strings.Replace(arg, ":", "", -1)] = newStr
			}
		}
	}
	return result
}
