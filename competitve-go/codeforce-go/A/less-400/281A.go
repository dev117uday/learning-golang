package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	istring, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	temp := strings.Split(istring, "")
	temp[0] = strings.ToUpper(temp[0])
	xstr := strings.Join(temp, "")
	fmt.Println(xstr)
}
