package main

import (
	"fmt"
)

func main() {
	sum := 1
	for sum < 10 {
		sum += sum
	}
	fmt.Println(sum)

	// infinite loop

	// for {
	// }

	sum = 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	sum = 1
	for sum < 1024 {
		sum += sum
	}
	fmt.Println(sum)

	pow := []int{1, 2, 4, 8, 16, 32, 64, 128}

	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

}
