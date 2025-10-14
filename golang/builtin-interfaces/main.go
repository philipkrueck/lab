package main

import (
	"fmt"
	"time"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

type CustomError struct {
	When time.Time
	What string
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("The error '%v' happened at %v", e.What, e.When)
}

func run() error {
	return &CustomError{
		time.Now(),
		"A really grave error.",
	}
}

func main() {
	philip := Person{"Philip", 26}
	fmt.Println(philip)

	if err := run(); err != nil {
		fmt.Println(err)
	}
}
