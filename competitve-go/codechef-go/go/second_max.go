package main

import "fmt"

func main() {

	col := 0
	fmt.Scanln(&col)
	var matrix = make([][3]int, col)
	for i := 0; i < col; i++ {
		for j := 0; j < 3; j++ {
			fmt.Scan(&matrix[i][j])
		}
	}
	for i := 0; i < col; i++ {
		fmt.Println(secondBiggest(matrix[i][0], matrix[i][1], matrix[i][2]))
	}

}

func secondBiggest(a, b, c int) int {
	if a > b && a > c {
		if b > c {
			return b
		} else {
			return c
		}
	}
	if b > a && b > c {
		if a > c {
			return a
		} else {
			return c
		}
	}
	if c > b && c > a {
		if a > b {
			return a
		} else {
			return b
		}
	}
	return 0
}
