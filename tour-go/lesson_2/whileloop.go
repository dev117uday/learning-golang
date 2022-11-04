package lesson_2

import (
	"fmt"
)

func while_loop() {
	sum := 1
	for sum < 1024 {
		sum += sum
	}
	fmt.Print(sum)
}
