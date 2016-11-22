//A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,

//a^2 + b^2 = c^2
//For example, 3^2 + 4^2 = 9 + 16 = 25 = 5^2.

//There exists exactly one Pythagorean triplet for which a + b + c = 1000.
//Find the product abc.

package main

import (
	"fmt"
)

func main() {

	for m := 2; m < 30; m++ {
		for n := 1; n < m; n++ {
			a := (m * m) - (n * n)
			b := 2 * m * n
			c := (m * m) + (n * n)
			if a+b+c == 1000 {
				fmt.Println(a, b, c, "==>", a+b+c)
			}
		}
	}

}
