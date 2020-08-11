package main

import (
	"fmt"
	"strings"
)

func main() {
	var num int32
	_, _ = fmt.Scanln(&num)
	var xstring string
	_, _ = fmt.Scanln(&xstring)
	temp := strings.Split(xstring,"")
	a := 0
	d := 0
	for i:=0;i<len(temp); i++ {
		if temp[i] == "A"{
			a++
		} else {
			d++
		}
	}
	if a>d {
		fmt.Println("Anton")
	} else if a<d {
		fmt.Println("Danik")
	} else {
		fmt.Println("Friendship")
	}
}