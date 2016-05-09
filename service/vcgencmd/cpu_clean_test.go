package vcgencmd

import "testing"

func TestClean(t *testing.T) {

	var tests = []struct {
		data     string
		expected string
	}{{
		data:     "temp=34.7'C",
		expected: "34.7",
	}}

	for _, test := range tests {
		result := Clean(test.data, "temp=", "'C")
		if test.expected != result {
			t.Error(test.expected, "!=", result)
		}
	}
}
