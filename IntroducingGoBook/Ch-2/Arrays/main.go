package main

import "fmt"

func main() {
	myArray := [5]float64{
		98,
		93,
		65,
		44,
		29,
	}
	var total float64

	for _, value := range myArray {
		total += value
	}
	fmt.Println(total / float64(len(myArray)))
}
