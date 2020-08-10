package main

import (
	"fmt"
	"math"
)

func funcPow( x,n,limit float64 ) float64 {

	if v:= math.Pow(x,n);
	v < limit {
		return v
	}
	return limit

}

func main(){	
	fmt.Print("3^2 < 20 == ",funcPow(3,2,20))		
}