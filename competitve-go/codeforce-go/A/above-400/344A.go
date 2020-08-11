package main

import (
	"fmt"
	"os"
)

func main() {
	var number int
	_, _ = fmt.Scanln(&number)
	var oned = make([]int, number)

	if number == 1 {
		_, _ = fmt.Scanln(&oned[0])
		fmt.Print("1")
		os.Exit(0)
	}
	count := 1
	_, _ = fmt.Scanln(&oned[0])
	for i:=1; i<number;i++ {
		_, _ = fmt.Scanln(&oned[i])
		if oned[i] != oned[i-1] {
			count++
		}
	}
	

	fmt.Println(count)


}
