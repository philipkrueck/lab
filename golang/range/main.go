package main

import "fmt"

func main() {
	var pow [10]int
	for i := range pow {
		pow[i] = 1 << i
	}

	for _, v := range pow {
		fmt.Printf("%d\n", v)
	}
}
