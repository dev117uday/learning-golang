package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	myReader,_ := bufio.NewReader(os.Stdin).ReadString('\n')
	fmt.Println("Input String : ",myReader)

	myNumber,_ := bufio.NewReader(os.Stdin).ReadString('\n')
	myInteger,_ := strconv.ParseFloat(strings.TrimSpace(myNumber),64)
	fmt.Println("my integer : ",myInteger)
}