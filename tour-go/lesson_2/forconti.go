package lesson_2

import (
	"fmt"
)

func for_loop() {
	sum := 1
	for sum < 10 {
		sum += sum
	}
	fmt.Print(sum)
}
