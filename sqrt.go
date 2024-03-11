package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.5
	prev := 0.0
	for a := 1; a < 10; a++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println(z, prev, math.Abs(z-prev))
		if math.Abs(z-prev) < 0.000000000000001 {
			return z
		}
		prev = z
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}
