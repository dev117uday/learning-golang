package main

import "fmt"

func main() {

	col := 0
	fmt.Scanln(&col)
	matrix := make([]int, col)
	for i := 0; i < col; i++ {
		fmt.Scan(&matrix[i])
		fmt.Println(matrix[i] * matrix[i])

	}

}
