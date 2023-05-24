package main

// Defer
// A defer statement defers the execution of a function until the surrounding function returns.
// The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.

import "fmt"

func printNo(number int) {
	defer fmt.Print(" ", number)
	if number == 1 {
		return
	}
	printNo(number - 1)
	return
}

func main() {
	defer fmt.Println("world")

	fmt.Println("hello")

	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	printNo(10)

	fmt.Println("done")
}
