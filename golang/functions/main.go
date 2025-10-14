package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	add := func(x, y float64) float64 {
		return x + y
	}

	hypothenuse := func(x, y float64) float64 {
		return math.Sqrt((x*x + y*y))
	}

	fmt.Println("simple sum", compute(add))
	fmt.Println("power", compute(math.Pow))
	fmt.Println("hypothenuse", compute(hypothenuse))
}
