package problems

import (
    "fmt"
    "sort"
)

// Multiples of 3 and 5
//
// If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.
// Find the sum of all the multiples of 3 or 5 below 1000.

func P1() string {

    var multiples []int

    // Multiples of 3
    for m := 0 ; m + 3 < 1000; {
        m += 3
        multiples = append(multiples, m)
    }

    // Multiples of 5
    for m := 0 ; m + 5 < 1000; {
        m += 5
        multiples = append(multiples, m)
    }

    // Sort for deduping durring summation
    sort.Sort(sort.IntSlice(multiples))

    // Sum multiples
    sum := 0
    for i, m := range multiples {
        // Skip if duplicate
        if i == 0 || m != multiples[i - 1] {
            sum += m
        } else {
            continue
        }
    }

    return fmt.Sprint(sum)

}
