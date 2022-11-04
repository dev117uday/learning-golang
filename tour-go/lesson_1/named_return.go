package lesson_1

import (
	"fmt"
)

func add(x, y int) (p int) {
	p = x + y
	return
}

func named_return() {

	fmt.Println("Sum is : ", add(3, 4))

}
