package main

import (
	"fmt"
	"strings"
)

func main() {
	var istring string
	_, _ = fmt.Scan(&istring)
	var check = []string{"h","e","l","l","o"}
	j:=0
	count :=0
	temp := strings.Split(istring,"")
	for i:=0; i<len(temp); i++ {
		if temp[i] == check[j] {
			count++
			if count == 5 {
				break
			} else {
				j++
			}
		}
	}
	if count == 5 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
