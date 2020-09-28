package main

import "fmt"

func main() {

	number := 0
	for {
		fmt.Scanln(&number)
		if number == 42 {
			break
		}
		fmt.Println(number)
	}

}
