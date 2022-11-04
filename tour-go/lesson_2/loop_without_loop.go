package lesson_2

import "fmt"

func printNo2(number int) {

	if number == 1 {
		fmt.Print(number)
		return
	}
	printNo(number - 1)
	fmt.Print(" ", number)
	return
}

func recusion() {

	var number int
	fmt.Scanln(&number)
	printNo2(number)

}
