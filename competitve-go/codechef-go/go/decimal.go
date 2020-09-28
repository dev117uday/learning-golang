package main

import "fmt"

func main() {

	x := 0
	fmt.Scanln(&x)
	var n, input int

	for i := 0; i < x; i++ {
		fmt.Scanln(&input)
		n = 1
		for n*2 <= input {
			n = n * 2
		}
		for n >= 1 {
			fmt.Print(input / n)
			input = input % n
			n = n / 2
		}
		fmt.Println("")
	}
}
