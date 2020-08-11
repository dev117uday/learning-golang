package main

import (
	"fmt"
	"os"
)

func main() {

	var number string
	count := 0
	_, _ = fmt.Scanln(&number)

	if number == "7" || number == "4" {
		fmt.Print("NO")
		os.Exit(0)
	}

	for i:=0; i<len(number); i++ {
		if string(number[i]) == "4" || string(number[i]) == "7" {
			continue
		} else {
			count++
			break
		}
	}

	if count == 0 {
		if len(number) == 4 || len(number) == 7 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	} else {

		counter := 0
		for i:=0; i<len(number); i++ {
			if string(number[i]) == "4" || string(number[i]) == "7" {
				counter++
			}
		}

		if counter == 4 || counter == 7 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}

	}

}
