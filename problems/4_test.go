package problems

import "testing"

func TestP4(t *testing.T) {
    a := P4()
    e := "906609"
    if a != e {
        t.Error("Answer given was", a, "expected", e)
    }
}