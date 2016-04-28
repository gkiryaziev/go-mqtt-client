package meminfo

import "testing"

func TestClean(t *testing.T) {
	in := "MemTotal:         996784 kB\n" +
		"MemFree:          918592 kB\n" +
		"MemAvailable:     938892 kB\n" +
		"Buffers:            7364 kB\n" +
		"Cached:            37604 kB\n"

	result := Clean(in, "MemTotal:", "MemFree:", "MemAvailable:")

	if result == nil {
		t.Error("Error, result = nil.")
	}

	memTotal := result["MemTotal"]
	memFree := result["MemFree"]
	memAvailable := result["MemAvailable"]

	if memTotal != "996784" || memFree != "918592" || memAvailable != "938892" {
		t.Errorf("Error, MemTotal = %s, MemFree = %s, MemAvailable = %s", memTotal, memFree, memAvailable)
	}
}
