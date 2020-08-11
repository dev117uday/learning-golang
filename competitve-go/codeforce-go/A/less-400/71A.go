package main

import (
	"fmt"
)

func main() {
	var number uint8
	_, _ = fmt.Scanln(&number)

	for number > 0 {
		var iString string
		_, _ = fmt.Scanln(&iString)
		if len(iString) <= 10 {
			fmt.Printf("%s\n",iString)
		} else {
			length := len(iString)
			fmt.Printf("%c%d%c\n", iString[0], length-2, iString[length-1])
		}
		number--
	}
}
