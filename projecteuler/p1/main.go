package main

import "fmt"

func main() {
	all := removeDuplicates(findmultiples(999, 3, 5))
	var total int
	for _, value := range all {
		total += value
	}
	fmt.Println(total)
}

func findmultiples(max int, dividers ...int) []int {
	var vals []int
	for _, divider := range dividers {
		for i := 2; i <= max; i++ {
			if i%divider == 0 {

				vals = append(vals, i)

			}
		}
	}

	return vals
}

func removeDuplicates(elements []int) []int {
	// Use map to record duplicates as we find them.
	encountered := map[int]bool{}
	result := []int{}

	for v := range elements {
		if encountered[elements[v]] == true {
			// Do not add duplicate.
		} else {
			// Record this element as an encountered element.
			encountered[elements[v]] = true
			// Append to result slice.
			result = append(result, elements[v])
		}
	}
	// Return the new slice.
	return result
}