package problems

import (
    "fmt"
    "math"
)

// Largest prime factor
//
// The prime factors of 13195 are 5, 7, 13 and 29.
// What is the largest prime factor of the number 600851475143?

func P3() string {

    n := 600851475143
    primes := eratosthenesSieve(int(math.Sqrt(float64(n))))

    for i := len(primes) - 1; ; i-- {
        if math.Mod(float64(n), float64(primes[i])) == 0 {
            return fmt.Sprint(primes[i])
        } else {
            continue
        }
    }

    return fmt.Sprint(primes)

}

func eratosthenesSieve(limit int) []int {
    ints := make(map[int]bool)

    // Generate slice of integers from 2 to limit
    for i := 2; i <= limit; i++ {
        ints[i] = true
    }

    // Remove composites
    for i := 2; i <= limit; i++ {
        // true == prime
        if ints[i] == true {
            // mark multiples as false
            for m := i + i; m <= limit; m += i {
                ints[m] = false
            }
        // false == composite, go to next item
        } else {
            continue
        }
    }

    var primes []int
    for i := 2; i <= limit; i++ {
        if ints[i] {
            primes = append(primes, i)
        }
    }

    return primes

}
