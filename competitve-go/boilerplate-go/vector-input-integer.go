package main

import "fmt"

func main() {

	var number int
	_, _ = fmt.Scanln(&number)
	var vector = make([]int, number)
	for i := 0; i < number; i++ {
		_, _ = fmt.Scan(&vector[i])
	}
	fmt.Println(vector)
}
