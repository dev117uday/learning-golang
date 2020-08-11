package main

import "fmt"

func main() {
	var num int
	_, _ = fmt.Scanln(&num)
	matrix := make([][2]int32, num)
	count := 0
	for i := 0; i < num; i++ {
		for j := 0; j < 2; j++ {
			_, _ = fmt.Scan(&matrix[i][j])
		}
		if matrix[i][0] < matrix[i][1] {
			if matrix[i][1]-matrix[i][0] >= 2 {
				count++
			}
		}
	}
	fmt.Print(count)
}
