package problems

import "testing"

func TestP3(t *testing.T) {
    a := P3()
    e := "6857"
    if a != e {
        t.Error("Answer given was", a, "expected", e)
    }
}