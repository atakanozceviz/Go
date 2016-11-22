package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	fmt.Println(largestPalindrome(100, 999))
}

func reverse(number int) int {

	value := strconv.Itoa(number)
	// Convert string to rune slice.
	data := []rune(value)
	result := []rune{}

	// Add runes in reverse order.
	for i := len(data) - 1; i >= 0; i-- {
		result = append(result, data[i])
	}

	i, _ := strconv.Atoi(string(result))

	return i
}


func largestPalindrome(min, max int) int {
	var numbers []int

	for i := max; i >= min; i-- {
		for j := max; j >= i; j-- {
			numbers = append(numbers, i*j)
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(numbers)))

	for _, value := range numbers {
		if value == reverse(value) {
			return value
		}

	}
	return 0
}
