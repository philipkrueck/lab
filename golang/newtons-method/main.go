package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	prev := -1.0

	for i := 0; i < 10; i++ {
		if math.Abs(z-prev) < 0.0001 {
			break
		}
		prev = z
		z -= (z*z - x) / (2 * z)
	}

	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
