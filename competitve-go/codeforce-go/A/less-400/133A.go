package main

import (
	"fmt"
	"strings"
)

func main() {
	var istring string
	_, _ = fmt.Scanln(&istring)
	temp :=  strings.Split(istring,"")
	flag := 0
	for i:=0; i<len(temp); i++{
		if temp[i] == "H" || temp[i] == "Q" || temp[i] == "9" {
			fmt.Println("YES")
			flag = 1
			break
		}
	}
	if flag == 0 {
		fmt.Println("NO")
	}
}
