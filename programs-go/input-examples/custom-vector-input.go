package main

import (
	"fmt"
)

func main() {
	numberOfStrings := 0
	var vector []string
	_, _ = fmt.Scanln(&numberOfStrings)

	{
		var number string
		for i := 0; i < numberOfStrings; i++ {
			_, _ = fmt.Scanln(&number)
			vector = append(vector, number)
		}
	}

}
