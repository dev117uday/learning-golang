package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }

func main() {
	var number, i uint8
	_, _ = fmt.Scanln(&number)
	number++
	var oned = make([]uint8, number, number+3)
	for i = 1; i < number; i++ {
		var p uint8
		scanf("%d", &p)
		oned[p] = i
	}

	fmt.Print(oned[1])
	for i := 2; i < len(oned); i++ {
		fmt.Print(" ", oned[i])
	}
	fmt.Print("\n")

}
