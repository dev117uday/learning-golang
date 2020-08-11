package main

import (
	"fmt"
)

func main() {
	var numberOfIntegers uint16
	var vector []uint32
	_, _ = fmt.Scanln(&numberOfIntegers)

	{
		var i uint16 = 0
		var number uint32
		for i = 0 ; i< numberOfIntegers; i++ {
			_, _ = fmt.Scanln(&number)
			vector = append(vector,number)
		}
	}

	for _,value := range vector {
		var temp uint32 = 0
		for value > 0 {
			temp = temp*10 + value%10
			value = value/10
		}
		fmt.Println(temp)
	}

}
