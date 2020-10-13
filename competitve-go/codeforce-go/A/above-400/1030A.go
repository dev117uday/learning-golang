package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

func scanf(f string, a ...interface{}) { fmt.Fscanf(reader, f, a...) }

func main() {

	var number, flag uint = 0, 0
	_, _ = fmt.Scanln(&number)
	var oned = make([]int, number, number+2)
	for i := 0; i < len(oned); i++ {
		scanf("%d", &oned[i])
		if oned[i] == 1 {
			flag = 1
		}
	}
	if flag == 0 {
		fmt.Print("EASY")
	} else {
		fmt.Print("HARD")
	}

}
