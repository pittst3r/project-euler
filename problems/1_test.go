package problems

import "testing"

func TestP1(t *testing.T) {
    a := P1()
    e := "233168"
    if a != e {
        t.Error("Answer given was", a, "expected", e)
    }
}