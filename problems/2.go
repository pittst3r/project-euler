package problems

import (
    "fmt"
    "math"
)

// Even Fibonacci numbers
//
// Each new term in the Fibonacci sequence is generated by adding the previous two terms. By starting with 1 and 2, the first 10 terms will be:
// 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...
// By considering the terms in the Fibonacci sequence whose values do not exceed four million, find the sum of the even-valued terms.

func P2() string {

    seq := []int{0}

    // Calculate sequence of fibs under 4mm
    for next := 1; next < 4000000; {
        seq = append(seq, next)
        next = seq[len(seq) - 1] + seq[len(seq) - 2]
    }

    // Sum even numbers
    sum := 0
    for _, n := range seq[1:] {
        if math.Mod(float64(n), 2) == 0 {
            sum += n
        }
    }

    return fmt.Sprint(sum)

}
