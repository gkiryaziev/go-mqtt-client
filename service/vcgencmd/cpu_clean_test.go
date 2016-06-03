package vcgencmd

import "testing"

func TestClean(t *testing.T) {

	var tests = []struct {
		data     string
		garbage  []string
		expected string
	}{{
		data:     "temp=34.7'C",
		garbage:  []string{"temp=", "'C"},
		expected: "34.7",
	}, {
		data:     "volt=1.2V",
		garbage:  []string{"volt=", "V"},
		expected: "1.2",
	}}

	for _, test := range tests {
		result := Clean(test.data, test.garbage...)
		if test.expected != result {
			t.Error(test.expected, "!=", result)
		}
	}
}
