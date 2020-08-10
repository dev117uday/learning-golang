// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// func main() {
// 	var vector []int64
// 	for {
// 		fmt.Print("Enter the number : ")
// 		reader := bufio.NewReader(os.Stdin)
// 		myRating, _ := reader.ReadString('\n')
// 		number := strings.TrimSpace(myRating)
// 		if number == string('x') || number == string('X') {
// 			for i := len(vector); i > 0; i-- {
// 				for j := 1; j < i; j++ {
// 					if vector[j-1] > vector[j] {
// 						intermediate := vector[j]
// 						vector[j] = vector[j-1]
// 						vector[j-1] = intermediate
// 					}
// 				}
// 			}
// 			fmt.Println(vector)
// 			os.Exit(0)
// 		} else {
// 			myInteger, _ := strconv.ParseInt(strings.TrimSpace(myRating), 10, 64)
// 			vector = append(vector, myInteger)
// 		}
// 	}
// }

package main

import "fmt" 

for input.Scan() 
    if input.text() == "X" {break}
    fmt.Println(input.text())
	
	fmt.Println("Val i %d\n", i)
	if i > 3 {
	   fmt.Printf("Extend slice")
	   sli = sli[:1]
	}
	j := i
	intVal, err := strconv.Atoi(input.text())
	if err != nil {
       fmt.Println(err)
   }
  
   sli[j] = intVal
	
  i += 1
}