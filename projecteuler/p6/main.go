package main

import "fmt"

func main() {
	limit := 100
	sumSq := 0
	sum := 0
	for i := 1; i <= limit; i++ {
		sum = sum + i
		sumSq = sumSq + i*i
	}
	fmt.Println(sum*sum - sumSq)
}
