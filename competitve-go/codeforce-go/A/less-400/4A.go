package main

import "fmt"

func main() {
	var number uint8
	_, _ = fmt.Scanln(&number)
	if number == 2 {
		fmt.Println("NO")
	} else if number%2 == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
