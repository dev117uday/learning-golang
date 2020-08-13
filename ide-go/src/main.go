package main

import (
	"fmt"
	"sort"
)

func main() {

	var number int
	_, _ = fmt.Scan(&number)
	var oned = make([]int,number)
	for i:=0; i<len(oned); i++ {
		_, _ = fmt.Scan(&oned[i])
	}
	sort.Ints(oned)
	counter := 0
	for i:=0; i<len(oned); i++ {
		for j:=0; j<len(oned); j++ {
			
		}
	}

	fmt.Println(counter)

}