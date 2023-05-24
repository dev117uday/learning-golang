package main

import (
	"fmt"
	"math"
)

func funcPow(x, n, limit float64) float64 {
	if v := math.Pow(x, n); v < limit {
		return v
	}
	return limit
}

func main() {
	grade := 90

	if grade > 90 {
		fmt.Print("Over 90")
	} else if grade <= 90 && grade > 80 {
		fmt.Print("passed")
	} else {
		fmt.Print("Fail")
	}

	fmt.Print("3^2 < 20 == ", funcPow(3, 2, 20))
}
