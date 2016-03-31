package problems

import (
	"fmt"
	"math"
)

// Smallest multiple
//
// 2520 is the smallest number that can be divided by each of the numbers from
// 1 to 10 without any remainder.
//
// What is the smallest positive number that is evenly divisible by all of the
// numbers from 1 to 20?

func P5() string {
	var multiple int

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	largestNumber := numbers[len(numbers)-1]

	// Loop through multiples of largest number
Loop:
	for multiplier := 2; ; multiplier++ {
		multiple = largestNumber * multiplier

		for ordinal := len(numbers) - 1; ; ordinal-- {
			mult := float64(multiple)
			number := float64(numbers[ordinal])
			if math.Mod(mult, number) != 0 {
				break
			} else if ordinal == 0 {
				break Loop
			}
		}
	}

	return fmt.Sprint(multiple)
}
