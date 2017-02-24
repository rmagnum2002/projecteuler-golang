package main

import (
	"fmt"
)

func main() {
	var number = 600851475143
	var primes []int
	for i := 2; i < 10000; i++ {
		if number%i == 0 {
			// fmt.Println(i)
			primes = append(primes, i)
		}
	}
	fmt.Println(primes[len(primes)-1])
}

// Largest prime factor
// Problem 3
// The prime factors of 13195 are 5, 7, 13 and 29.
// What is the largest prime factor of the number 600851475143 ?
// => 6857
