package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

func scanf(f string, a ...interface{}) { fmt.Fscanf(reader, f, a...) }

func main() {

	var number, height, i int16 = 0, 0, 0
	_, _ = fmt.Scan(&number)
	_, _ = fmt.Scan(&height)

	var oned = make([]int16, number, number+4)

	for i = 0; i < number; i++ {
		scanf("%d ", &oned[i])
	}

	var count int = 0
	for i = 0; i < number; i++ {
		for oned[i] > 0 {
			count++
			oned[i] = oned[i] - height
		}
	}
	fmt.Print(count)

}
