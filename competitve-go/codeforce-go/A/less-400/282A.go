package main

import (
	"fmt"
)

func main() {
	var bitland []string
	number := 0
	_, _ = fmt.Scanln(&number)
	for number > 0 {
		var istring string
		_, _ = fmt.Scanln(&istring)
		bitland = append(bitland, istring)
		number--
	}
	var count int32 = 0

	for i := 0; i < len(bitland); i++ {
		if string(bitland[i]) == "++X" {
			count++
		} else if string(bitland[i]) == "X++" {
			count++
		} else {
			count--
		}
	}

	fmt.Println(count)
}
