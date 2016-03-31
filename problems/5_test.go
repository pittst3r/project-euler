package problems

import "testing"

func TestP5(t *testing.T) {
	a := P5()
	e := "2520"
	if a != e {
		t.Error("Answer given was", a, "expected", e)
	}
}
