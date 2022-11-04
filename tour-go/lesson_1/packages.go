package lesson_1

import (
	"fmt"
	"math"
	"math/rand"
)

func packages() {
	fmt.Println("My favorite number is", rand.Intn(10))
	// this is always output 1
	fmt.Println("Sqaure Root of 5 is : ", math.Sqrt(5))

	fmt.Println("Math pie is : ", math.Pi)
}
