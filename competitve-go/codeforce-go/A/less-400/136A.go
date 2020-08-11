package main

import "fmt"

func main() {
	var number,i uint8
	_, _ = fmt.Scanln(&number)
	number++
	var oned = make([]uint8,number)
	for i=1; i<number; i++ {
		var p uint8
		_, _ = fmt.Scan(&p)
		oned[p] = i
	}
	fmt.Print(oned[1])
	for i:=2;i<len(oned);i++ {
		fmt.Print(" ",oned[i])
	}
	fmt.Print("\n")


}
