package main

import "fmt"

const small = 1

func needFloat(x float64) float64 { return x * 2 }

func main() {
	fmt.Printf("small has type '%T' and value '%v'\n", small, small)

	// the constant now takes the required type of context (float64)
	x := needFloat(small)
	fmt.Printf("x has type '%T' and value '%v'\n", x, x)
}
