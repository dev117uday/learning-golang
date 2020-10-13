package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

func scanf(f string, a ...interface{}) { fmt.Fscanf(reader, f, a...) }

func main() {
	var number, i int = 0, 1
	_, _ = fmt.Scanln(&number)
	var oned = make([]int, number, number+3)

	if number == 1 {
		_, _ = fmt.Scanln(&oned[0])
		fmt.Print("1")
		os.Exit(0)
	}
	count := 1
	_, _ = fmt.Scanln(&oned[0])
	for i = 1; i < number; i++ {
		scanf("%d\n", &oned[i])
		if oned[i] != oned[i-1] {
			count++
		}
	}

	fmt.Println(count)

}
