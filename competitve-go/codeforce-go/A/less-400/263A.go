package main

import (
	"fmt"
	"math"
)

func main() {
	var matrix [5][5]int32
	var m, n = 0, 0

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			_, _ = fmt.Scan(&matrix[i][j])
			if matrix[i][j] == 1 {
				m = i
				n = j
			}
		}
	}

	temp := math.Abs(float64(m)-2) + math.Abs(float64(n)-2)
	fmt.Println(temp)
}
