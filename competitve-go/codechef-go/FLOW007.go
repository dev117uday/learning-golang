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
	var numberOfIntegers uint16

	scanf("%d\n", &numberOfIntegers)
	vector := make([]uint32, numberOfIntegers, numberOfIntegers+3)

	{
		var i uint16 = 0
		for i = 0; i < numberOfIntegers; i++ {
			scanf("%d\n", &vector[i])
		}
	}

	for _, value := range vector {
		var temp uint32 = 0
		for value > 0 {
			temp = temp*10 + value%10
			value = value / 10
		}
		fmt.Println(temp)
	}

}
