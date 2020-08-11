package main

import (
	"fmt"
)

func main() {

	var vector []int8
	var input int8
	for {
		fmt.Scanln(&input)
		if input == int8(42) {
			for _,value := range vector {
				fmt.Println(value)
			}
			break
		} 
		vector = append(vector, input)
	}

}