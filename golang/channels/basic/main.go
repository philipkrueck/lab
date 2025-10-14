package main

import "fmt"

func sum(elements []int, c chan int) {
	sum := 0
	for _, el := range elements {
		sum += el
	}
	c <- sum
}

func main() {
	elements := []int{1, 2, 3, 4, 5, 6}

	c := make(chan int)
	go sum(elements[:len(elements)/2], c)
	go sum(elements[len(elements)/2:], c)

	x, y := <-c, <-c

	fmt.Println(x, y, x+y)

	cb := make(chan int, 2)
	cb <- 1
	cb <- 2
	r := <-cb
	cb <- 3

	fmt.Println(<-cb, <-cb, r)
}
