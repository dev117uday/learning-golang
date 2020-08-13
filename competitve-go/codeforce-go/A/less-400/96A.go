package main

import (
	"fmt"
	"strings"
)

func main() {
	var istring string
	fmt.Scan(&istring)
	temp := strings.Split(istring,"")
	answer := "NO"
	for i:=0;i<len(temp)-6; i++ {
		t := temp[i]
		if temp[i+1] == t && temp[i+2] == t && temp[i+3] == t && temp[i+4] == t && temp[i+5] == t && temp[i+6] == t {
			answer = "YES"
			break
		}
	}
	fmt.Println(answer)
}
