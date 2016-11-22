package main

import "math"
import "fmt"

func main() {
	sum := 0
	for i := 0; i < 2000000; i++ {
		if isPrime(i) {
			sum += i
		}
	}
	fmt.Println(sum)
}

func isPrime(n int) bool {

	if n == 1 {
		return false
	} else if n < 4 {
		return true
	} else if n%2 == 0 {
		return false
	} else if n < 9 {
		return true
	} else if n%3 == 0 {
		return false
	}
	r := math.Floor(math.Sqrt(float64(n)))

	f := 5
	for f <= int(r) {
		if n%f == 0 {
			return false
		}
		if math.Mod(float64(n), float64(f+2)) == 0 {
			return false
		}
		f = f + 6
	}

	return true

}
