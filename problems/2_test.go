package problems

import "testing"

func TestP2(t *testing.T) {
    a := P2()
    e := "4613732"
    if a != e {
        t.Error("Answer given was", a, "expected", e)
    }
}