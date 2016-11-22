package main

import (
	"fmt"
	"math"
)

func main() {

	var total float64
	for _, value := range fibo(4000000) {
		if math.Mod(value, 2) == 0 {
			total += value
		}
	}
	fmt.Println(total)
}

func fibo(max float64) []float64 {

	f := []float64{1, 2}
	i := 2

	for {
		var nextNumber float64
		nextNumber = f[i-1] + f[i-2]

		if nextNumber <= max {
			f = append(f, nextNumber)
		} else {
			break
		}
		i++

	}
	return f
}
