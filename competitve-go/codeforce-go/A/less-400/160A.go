package main

import (
	"fmt"
	"sort"
)

func main() {
	var num int
	_, _ = fmt.Scan(&num)
	var oned = make([]int,num)
	sum := 0
	for i:=0; i<num; i++ {
		fmt.Scan(&oned[i])
		sum = sum + oned[i]
	}
	sort.Sort(sort.Reverse(sort.IntSlice(oned)))
	temp :=0
	count := 0
	for i:=0; i<len(oned); i++ {
		if temp > sum {
			break
		}
		temp = temp+oned[i]
		count++
		sum = sum - oned[i]
	}
	fmt.Println(count)

}
