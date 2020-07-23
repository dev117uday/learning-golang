package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func findian(xstr string) string {
	xstr = strings.ToLower(xstr)
	b := strings.ContainsAny(xstr, "a")

	if xstr[0] == 'i' && xstr[len(xstr)-2] == 'n' && b {
		return "Found!"
	}else{
		return "Not Found!"
	}
}

func main(){
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter a string: ")
	inputString ,_ := reader.ReadString('\n')
	fmt.Println(findian(inputString))

}