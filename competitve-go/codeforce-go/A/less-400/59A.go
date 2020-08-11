package main

import (
	"fmt"
	"strings"
)

func main() {
	var word string
	_, _ = fmt.Scanln(&word)

	lower := 0
	upper := 0

	for i:=0; i<len(word); i++ {
		if word[i] >= 65 && word[i] <= 90 {
			upper ++
		} else {
			lower++
		}
	}
	if upper > lower {
		fmt.Println(strings.ToUpper(word))
	} else  {
		fmt.Println(strings.ToLower(word))
	}
}