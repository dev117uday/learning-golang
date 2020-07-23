package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var vector []int64
	{
		reader := bufio.NewReader(os.Stdin)
		integerString,_ := reader.ReadString('\n')
		tempArray := strings.Split(integerString," ")
		for i:=0; i<len(tempArray); i++ {
			temp,_ := strconv.ParseInt(strings.TrimSpace(tempArray[i]),10,64)
			vector = append(vector,temp)
		}
	}
	fmt.Println("\nVector : ",vector)
}