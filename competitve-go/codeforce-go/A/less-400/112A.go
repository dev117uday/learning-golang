package main

import (
	"fmt"
	"strings"
)

func main() {

	var string1 string
	var string2 string

	_, _ = fmt.Scanln(&string1)
	_, _ = fmt.Scanln(&string2)

	str1 := strings.ToLower(string1)
	str2 := strings.ToLower(string2)

	for i := 0; i < len(str1); i++ {
		if str1[i] > str2[i] {
			fmt.Println("1")
			break
		} else if str1[i] < str2[i] {
			fmt.Println("-1")
			break
		} else if i == len(str1)-1 {
			fmt.Println("0")
			break
		}
	}

}
