package main

import (
	"fmt"
	"math"
	"math/rand"
)

func addTwoInt(x int, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x - y
}

func swap(x, y string) (string, string) {
	return y, x
}

func swap_nums() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}

func add(x, y int) (p int) {
	p = x + y
	return
}

func main() {
	fmt.Println("Sum of 4 and 6 is : ", addTwoInt(4, 6))
	fmt.Println("Difference of 6 and 4 is :", subtract(6, 4))
	swap_nums()
	fmt.Println("Sum is : ", add(3, 4))

	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Println("Sqaure Root of 5 is : ", math.Sqrt(5))
	fmt.Println("Math pie is : ", math.Pi)
}
