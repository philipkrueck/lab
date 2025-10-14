package main

import "fmt"

type Abser interface {
	Abs() float64
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func typeSwitch(unknown interface{}) {
	switch v := unknown.(type) {
	case int:
		fmt.Printf("Twice %d is %d.\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long.\n", v, len(v))
	default:
		fmt.Println("Default")
	}
}

func typeAssertion(unknown interface{}) {
	v, ok := unknown.(int)
	fmt.Println("v:", v, "is of type int:", ok)

	s, ok := unknown.(string)
	fmt.Println("s:", s, "is of type string:", ok)

	s = unknown.(string) // panic
	fmt.Println("s:", s, "Above line will panic")
}

func main() {
	var a Abser
	f := MyFloat(6)
	a = f
	fmt.Println("abs", a.Abs())

	var unknown interface{} = 42

	typeSwitch(unknown)
	typeAssertion(unknown)
}
