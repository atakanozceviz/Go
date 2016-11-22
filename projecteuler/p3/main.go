package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(findBigestPrime(600851475143))
}

func findBigestPrime(number float64) float64 {
	var i float64
	for i = 2; i < number; i++ {
		if math.Mod(number, i) == 0 {
			return findBigestPrime(number / i)
		}
	}
	return number
}
