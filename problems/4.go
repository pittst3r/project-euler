package problems

import (
    "fmt"
    "strings"
    "strconv"
    "sort"
    "math"
)

// Largest palindrome product
//
// A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.
// Find the largest palindrome made from the product of two 3-digit numbers.

func P4() string {

    var palindromes []int64

    // Calculate palindromes between 10000 and 998001
    for n := int64(100); ; n++ {
        a, b := genPalindromes(n)
        if a <= 998001 {
            palindromes = append(palindromes, a)
        }
        if b <= 998001 {
            palindromes = append(palindromes, b)
        }
        if a > 998001 && b > 998001 {
            break
        }
    }

    // Sort our palindromes
    sort.Sort(sort.Reverse(Int64Slice(palindromes)))

    var winner int64

    // Trial divide each palindrome in descending order starting from its square root going up
    PalLoop:
        for _, p := range palindromes {
            sqrt := int64(math.Sqrt(float64(p)))

            if i64Size(sqrt) == 3 {
                if sqrt * sqrt == p {
                    winner = p
                    break PalLoop
                }
                for i := sqrt + 1; i64Size(i) == 3; i++ {
                    if quo := p / i; divisible(p, i) && i64Size(quo) == 3 {
                        winner = p
                        break PalLoop
                    }
                }
            }
        }

    return fmt.Sprint(winner)

}

// Generates two palindromes using the 'seed':
//   1. Palindrome using last digit as middle digit in odd-digited integer (e.g. 105 -> 10501)
//   2. Palindrome that is simple mirror image of seed (e.g. 105 -> 105501)
func genPalindromes(seed int64) (odd, even int64) {

    // Convert to string
    s := fmt.Sprint(seed)
    sSlice := strings.Split(s, "")

    // Palindrome with odd number of digits
    firsts := sSlice[:i64Size(seed) - 1]
    middle := sSlice[i64Size(seed) - 1]
    lasts := reverse(firsts)
    oSlice := make([]string, (len(firsts) + 1 + len(lasts)))
    copy(oSlice[:len(firsts)], firsts)
    oSlice[len(firsts)] = middle
    copy(oSlice[len(firsts) + 1:], lasts)
    odd, _ = strconv.ParseInt(strings.Join(oSlice, ""), 0, 64)

    // Palindrome with even number of digits
    eSlice := make([]string, (i64Size(seed) * 2))
    copy(eSlice[:len(sSlice)], sSlice)
    copy(eSlice[len(sSlice):len(eSlice)], reverse(sSlice))
    even, _ = strconv.ParseInt(strings.Join(eSlice, ""), 0, 64)

    return
}

func reverse(s []string) []string {
    if len(s) == 1 {
        return s
    }
    trans := make([]string, len(s))
    for i, e := (len(s) - 1), 0; i >= 0; i, e = (i - 1), (e + 1) {
        trans[e] = s[i]
    }
    return trans
}

func i64Size(i int64) int {
    return len(strings.Split(fmt.Sprint(i), ""))
}

func divisible(i, d int64) bool {
    return math.Mod(float64(i), float64(d)) == 0
}

type Int64Slice []int64
func (p Int64Slice) Len() int           { return len(p) }
func (p Int64Slice) Less(i, j int) bool { return p[i] < p[j] }
func (p Int64Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
