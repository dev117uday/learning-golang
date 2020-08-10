package main

import(
	"fmt"
)

func main(){
	grade := 90
	if grade >90 {
		fmt.Print("Over 90")
	} else if grade <=90 && grade >80 {
		fmt.Print("passed")
	} else {
		fmt.Print("Fail")
	}

}