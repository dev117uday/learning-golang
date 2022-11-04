package lesson_2

import "fmt"

func printNo(number int) {
	defer fmt.Print(" ", number)
	if number == 1 {
		return
	}
	printNo(number - 1)
	return
}

func loop_wo_loop() {
	printNo(10)
}
