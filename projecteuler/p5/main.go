package main

import (
	"fmt"
)

var m = 1

func main() {

	fmt.Println(initalize())

}

func initalize() int {
	for k := range primesUntilNumber(20) {
		m *= k
	}
	theEnd(numbersNotDivisibleUntil(m, 20))
	return m
}

func findBigestPrime(number int) int {

	for i := 2; i < number; i++ {
		if number%i == 0 {
			return findBigestPrime(number / i)
		}
	}
	return number
}

func primesUntilNumber(number int) map[int]string {
	primes := make(map[int]string)
	for i := 2; i < number; i++ {
		if _, ok := primes[findBigestPrime(i)]; !ok {
			primes[i] = ""
		}
	}
	return primes
}

func numbersNotDivisibleUntil(number, until int) []int {
	var notDivisible []int
	for i := 2; i <= until; i++ {
		if number%i != 0 {
			notDivisible = append(notDivisible, i)
		}
	}
	return notDivisible
}

func theEnd(notDivisible []int) int {
	for _, v := range notDivisible {
		if v%2 == 0 {
			m *= 2
			return theEnd(numbersNotDivisibleUntil(m, 20))
		} else if v%3 == 0 {
			m *= 3
			return theEnd(numbersNotDivisibleUntil(m, 20))
		} else if v%5 == 0 {
			m *= 5
			return theEnd(numbersNotDivisibleUntil(m, 20))
		}
	}
	return m
}
