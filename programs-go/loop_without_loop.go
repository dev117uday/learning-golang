package main

import "fmt"

func printNo(number int) {

	if number == 1 {
		fmt.Print(number)
		return
	}
	printNo(number - 1)
	fmt.Print(" ", number)
	return
}

func main() {

	var number int
	fmt.Scanln(&number)
	printNo(number)

}
