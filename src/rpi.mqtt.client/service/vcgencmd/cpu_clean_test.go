package vcgencmd

import "testing"

func TestClean(t *testing.T) {
	in := "temp=34.7'C"

	out := "34.7"

	result := Clean(in, "temp=", "'C")

	if result != out {
		t.Errorf("Error, expected %s but return %s", out, result)
	}
}
