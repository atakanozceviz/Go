package main

import (
	"fmt"
)

func main() {
	fmt.Println(primesUntilNumber(9999999))
}

func findBigestPrime(number int) int {

	for i := 2; i < number; i++ {
		if number%i == 0 {
			return findBigestPrime(number / i)
		}
	}
	return number
}

func primesUntilNumber(number int) int {
	var primes []int
	primesM := make(map[int]string)
	for i := 2; i < number; i++ {
		if len(primes) >= 10001 {
			break
		}
		if _, ok := primesM[findBigestPrime(i)]; !ok {
			primesM[i] = ""
			primes = append(primes, i)
		}
	}
	return primes[10000]
}
