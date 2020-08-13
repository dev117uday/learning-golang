package main

import "fmt"

func printNo(number int) {
	defer fmt.Print(" ", number)
	if number == 1 {
		return
	}
	printNo(number - 1)
	return
}

func main() {
	printNo(10)
}
