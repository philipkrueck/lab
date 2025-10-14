package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y + v.Y)
}

func (v Vertex) Scale(f float64) {
	v.X = f * v.X
	v.Y = f * v.Y
}

type MyFloat float64

func (f MyFloat) Abs() MyFloat {
	if f < 0 {
		return -f
	}
	return f
}

func main() {
	v := Vertex{1, 2}

	fmt.Println("distance", v.Abs())

	f := MyFloat(-math.Sqrt2)
	fmt.Println("f", f.Abs())

	v.Scale(10)
	fmt.Println("v scaled", v)
}
