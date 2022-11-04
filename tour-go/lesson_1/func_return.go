package lesson_1

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func swap_nums() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
