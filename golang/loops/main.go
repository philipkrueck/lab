package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// while loop
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	// inifinite loop
	for {
	}
}
