package main

import "fmt"

func main() {
	myArray := [5]float64{0, 1, 2, 3, 4}
	slice1 := myArray[0:3] //Using [low:high] to create slice
	fmt.Println(slice1)

	slice2 := []int{1, 2, 3}
	slice3 := append(slice2, 4, 5)
	fmt.Println(slice3, slice2)

	//slice4 := make([]float64, 5, 10) ==> A slice of length 5 with a capacity of 10
	slice4 := make([]int, 2)
	slice5 := []int{1, 2, 3}
	copy(slice4, slice5)
	fmt.Println(slice4, slice5)

}
