package main

import (
	"fmt"
	"os"
)

func main() {
	var string1 string
	var string2 string
	_, _ = fmt.Scanln(&string1)
	_, _ = fmt.Scanln(&string2)

	for i:=0; i<len(string2); i++ {
		temp := len(string2) - 1 - i
		if string1[i] != string2[temp]{
			fmt.Println("NO")
			os.Exit(0)
		}
	}
	fmt.Println("YES")
}