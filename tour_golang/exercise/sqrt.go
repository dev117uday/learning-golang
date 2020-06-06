package main

import (
	"fmt"
	"math"
)

func sqrtOfNum(x uint32) float64 {
	var temp float64 = float64(x)
	for i:=0; i<20; i++ {
		temp = 0.5*(temp + float64(x)/temp)
	}
	return temp
	 
}	

func main()  {

	var someVar uint32 = 0;
	fmt.Println("Enter the number : ")
	fmt.Scanf("%d", &someVar)
	fmt.Print("\n")
	fmt.Printf("SquareRoot using custom function : %.3f",sqrtOfNum(someVar))
	fmt.Print("\n")
	fmt.Printf("SquareRoot using in-built function : %.3f",math.Sqrt(float64(someVar)))
	fmt.Print("\n")
}