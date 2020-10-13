package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }

func main() {
	var num int32
	scanf("%d\n", &num)
	var xstring string
	scanf("%s\n", &xstring)
	temp := strings.Split(xstring, "")
	var a, d, i int = 0, 0, 0
	for i = 0; i < len(temp); i++ {
		if temp[i] == "A" {
			a++
		} else {
			d++
		}
	}
	if a > d {
		fmt.Println("Anton")
	} else if a < d {
		fmt.Println("Danik")
	} else {
		fmt.Println("Friendship")
	}
}
