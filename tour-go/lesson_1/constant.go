package lesson_1

import "fmt"

// const pi
const pi = 3.14

func hello_world() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
