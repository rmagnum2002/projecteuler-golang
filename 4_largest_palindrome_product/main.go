package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var min = 100
	var max = 1000
	candidates := []int{}

	for i := max; i >= min; i-- {
		for n := max; n >= min; n-- {
			var number int
			var nts string

			number = i * n
			nts = strconv.Itoa(number)
			if nts == Reverse(nts) {
				candidates = append(candidates, number)
			}
		}
	}
	sort.Ints(candidates)
	fmt.Println(candidates[len(candidates)-1])
}

func Reverse(s string) string {
	var reverse string
	for i := len(s) - 1; i >= 0; i-- {
		reverse += string(s[i])
	}
	return reverse
}

// Largest palindrome product
// Problem 4
// A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.
// Find the largest palindrome made from the product of two 3-digit numbers.
// => 906609
