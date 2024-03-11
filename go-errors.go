package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %f", float64(e))
}

func Sqrt(x float64) (float64, error) {
	z := 1.5
	prev := 0.0
	if x < 0 {
		return x, ErrNegativeSqrt(x)
	}
	for a := 1; a < 10; a++ {
		z -= (z*z - x) / (2 * z)
		//fmt.Println(z, prev, math.Abs(z-prev))
		if math.Abs(z-prev) < 0.000000000000001 {
			return z, nil
		}
		prev = z
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
