package main

import (
	"fmt"
)

func main() {

	var cols, rows = 0, 0
	fmt.Print("Enter the number of cols : ")
	_, _ = fmt.Scan(&cols)
	fmt.Print("Enter the number of rows : ")
	_, _ = fmt.Scan(&rows)

	var twodslices = make([][]int, rows)
	var i int
	for i = range twodslices {
		twodslices[i] = make([]int, cols)
	}

	for i := 0; i < rows; i++ {
		for j:=0; j<cols; j++ {
			fmt.Scan(&twodslices[i][j])
		}
	}
	fmt.Println(twodslices)
}