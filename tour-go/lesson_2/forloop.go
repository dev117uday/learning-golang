package lesson_2

import (
	"fmt"
)

func for_loop_ex() {

	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Print(sum)

}
