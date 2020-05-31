package main

import(
	"fmt"
)

func add(x , y int) (p int) {
	p = x+y
	return
}

func main(){

	fmt.Println("Sum is : ",add(3,4))

}