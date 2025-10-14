package main

import "fmt"

func fibonacci() func() int {
	x, y := 0, 1

	return func() (z int) { // function captures x & y
		z, x, y = x, y, x+y
		return
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
