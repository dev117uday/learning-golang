package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	fmt.Print("Enter your full name : ")
	myString, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	fmt.Print("myString : ",myString)
	fmt.Print("Enter your rating : ")
	reader := bufio.NewReader(os.Stdin)
	myRating,_ := reader.ReadString('\n')
	myInteger,_ := strconv.ParseFloat(strings.TrimSpace(myRating),64)
	fmt.Println("Rating entered : ",myInteger)
}