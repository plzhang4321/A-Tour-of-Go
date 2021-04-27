package main

import (
	"fmt"
	"math"
)

/* Newton
Zn+1 = Z - (Z^2+N)/2*Z
*/

func Sqrt(x float64) float64 {
	z := x / 2
	err := 1e-15
	if x < 0 {
		return z
	} else {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
			if err >= math.Abs(math.Sqrt(x)-z) {
				break
			} else {
				z -= (z*z - x) / (2 * z)
			}
		}
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
