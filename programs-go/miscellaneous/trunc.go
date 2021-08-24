package main

import (
	"fmt"
	"math"
)

func main(){

	var num float64 = 0.0;
	fmt.Printf("	Enter a float : ")
	fmt.Scanf("%f",&num)
	var xnum int64 = int64(math.Floor(num))
	fmt.Printf("	Converted int is : %d\n",xnum)

}