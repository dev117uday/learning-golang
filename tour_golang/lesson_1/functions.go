package main 

import (
	"fmt"
)

func addTwoInt (x int, y int) (int) {
	return x+y
}

func subtract(x , y int) (int){
	return x-y
}

func main(){
	fmt.Println("Sum of 4 and 6 is : ",addTwoInt(4,6))
	fmt.Println("Difference of 6 and 4 is :",subtract(6,4))
}