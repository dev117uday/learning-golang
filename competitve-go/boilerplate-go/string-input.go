package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Print("Enter your full name : ")
	myString, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	fmt.Print("myString : ", myString)

}
