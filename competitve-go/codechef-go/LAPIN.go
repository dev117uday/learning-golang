package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }

func main() {
	var numberOfIntegers uint8
	var vector []string
	scanf("%d\n", &numberOfIntegers)

	// input strings into a vector
	{
		var i uint8 = 0
		var number string
		for i = 0; i < numberOfIntegers; i++ {
			scanf("%s\n", &number)
			vector = append(vector, number)
		}
	}

	for _, value := range vector {
		length := len(value)
		if length%2 == 0 {

			first := value[length/2:]
			temp := strings.Split(first, "")
			sort.Strings(temp)
			first = strings.Join(temp, "")

			second := value[:length/2]
			temp = strings.Split(second, "")
			sort.Strings(temp)
			second = strings.Join(temp, "")

			if first == second {
				fmt.Println("YES")
			} else {
				fmt.Println("NO")
			}

		} else {

			first := value[:length/2]
			temp := strings.Split(first, "")
			sort.Strings(temp)
			first = strings.Join(temp, "")

			second := value[length/2+1:]
			temp = strings.Split(second, "")
			sort.Strings(temp)
			second = strings.Join(temp, "")

			if first == second {
				fmt.Println("YES")
			} else {
				fmt.Println("NO")
			}
		}
	}

}
