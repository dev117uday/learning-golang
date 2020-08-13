package main

import (
	"fmt"
	"strings"
)

func main() {
	var istring string
	_, _ = fmt.Scan(&istring)
	istring = strings.ToLower(istring)
	temp := strings.Split(istring,"")
	var tempArray = []string{"."}
	for i:=0; i<len(temp); i++ {
		if temp[i] == "a" || temp[i] == "e" || temp[i] == "i" || temp[i] == "u" || temp[i] == "o" || temp[i] == "y" {
			continue
		} else {
			tempArray = append(tempArray,temp[i])
			tempArray = append(tempArray,".")
		}
	}
	tempArray = tempArray[:len(tempArray)-1]
	istring  = strings.Join(tempArray,"")
	fmt.Println(istring)

}
