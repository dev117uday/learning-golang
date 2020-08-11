package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var istring string
	_, _ = fmt.Scanln(&istring)
	xstring := strings.Split(istring, "+")
	var numArray []int
	for i := 0; i < len(xstring); i++ {
		inumber, _ := strconv.Atoi(xstring[i])
		numArray = append(numArray, inumber)
	}
	sort.Ints(numArray)
	for i := 0; i < len(numArray); i++ {
		if i == len(numArray)-1 {
			fmt.Print(numArray[i])
		} else {
			fmt.Printf("%d+", numArray[i])
		}
	}
}
