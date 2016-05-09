package meminfo

import "testing"

func TestClean(t *testing.T) {

	var tests = []struct {
		data         string
		memTotal     string
		memFree      string
		memAvailable string
	}{{
		data: "MemTotal:         996784 kB\n" +
			"MemFree:          918592 kB\n" +
			"MemAvailable:     938892 kB\n",
		memTotal:     "996784",
		memFree:      "918592",
		memAvailable: "938892",
	}}

	for _, test := range tests {
		result := Clean(test.data, "MemTotal:", "MemFree:", "MemAvailable:")

		if result == nil {
			t.Error("Error, result = nil.")
		}

		if test.memTotal != result["MemTotal"] {
			t.Error(test.memTotal, "!=", result["MemTotal"])
		}

		if test.memFree != result["MemFree"] {
			t.Error(test.memFree, "!=", result["MemFree"])
		}

		if test.memAvailable != result["MemAvailable"] {
			t.Error(test.memAvailable, "!=", result["MemAvailable"])
		}
	}
}
