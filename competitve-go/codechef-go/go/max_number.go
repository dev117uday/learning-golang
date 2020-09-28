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
		fmt.Println(max(matrix[i][0], matrix[i][1], matrix[i][2]))
	}

}

func max(a, b, c int) int {
	if a > b && a > c {
		return a
	}
	if b > a && b > c {
		return b
	}
	if c > b && c > a {
		return c
	}
	return 0
}
