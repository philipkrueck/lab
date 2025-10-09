package main

import (
	"fmt"
	"strings"
)

func main() {
	creatingSlices()

	slicesOfSlices()

	appendSlices()
}

func creatingSlices() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	s := primes[:4]
	printSlice("s", s)

	t := make([]int, 0, 5)
	printSlice("t", t)

	u := t[0:cap(t)]
	printSlice("u", u)

	v := t[1:cap(u)]
	printSlice("v", v)
}

func printSlice(name string, slice []int) {
	fmt.Printf("%s: %v, len: %d, cap: %d\n", name, slice, len(slice), cap(slice))
}

func slicesOfSlices() {
	board := [][]string{
		{"x", "_", "O"},
		{"_", "x", "O"},
		{"_", "O", "x"},
	}

	fmt.Println("\nTik Tak Toe:")
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func appendSlices() {
	backing := [3]int{1, 2, 3}
	slice := backing[:]
	fmt.Println()
	fmt.Printf("slice: %v, backing: %v\n", slice, backing)

	slice[1] = 9 // will change the backing array
	fmt.Printf("slice: %v, backing: %v\n", slice, backing)

	slice = append(slice, 4) // now will resize backing array

	slice[1] = 7 // should only affect the slice, not the backing array
	fmt.Printf("slice: %v, backing: %v\n", slice, backing)
}
