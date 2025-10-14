package main

import "fmt"

// Generic Functions
func Index[T comparable](s []T, x T) int {
	for i, el := range s {
		if el == x {
			return i
		}
	}
	return -1
}

// Generic Types
type List[T any] struct {
	val  T
	next *List[T]
}

func main() {
	// Generic Functions
	si := []int{1, 2, 3, 4}
	eli := 4
	fmt.Println("index of", eli, "in slice", si, "is", Index(si, eli))

	ss := []string{"foo", "bar", "baz"}
	els := "hello"
	fmt.Println("index of", els, "in slice", ss, "is", Index(ss, els))

	// Generic Types
	ll := List[int]{1, &List[int]{2, nil}}
	fmt.Println("ll:", ll.val, "->", ll.next.val)
}
