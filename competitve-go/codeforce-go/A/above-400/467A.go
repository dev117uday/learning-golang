package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

func scanf(f string, a ...interface{}) { fmt.Fscanf(reader, f, a...) }

func main() {
	var num uint
	_, _ = fmt.Scanln(&num)
	matrix := make([][2]uint32, num, num+3)
	var count, i uint = 0, 0
	for i = 0; i < num; i++ {
		scanf("%d %d\n", &matrix[i][0], &matrix[i][1])
		if matrix[i][0] < matrix[i][1] {
			if matrix[i][1]-matrix[i][0] >= 2 {
				count++
			}
		}
	}
	fmt.Print(count)
}
