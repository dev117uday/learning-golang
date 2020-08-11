package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var number int32
	_, _ = fmt.Scanln(&number)
	istring, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	xchar := istring[0]
	count := 0

	for i := 1; i < len(istring); i++ {
		if istring[i] == xchar {
			count++
		}
		xchar = istring[i]
	}

	fmt.Println(count)

}
