package lesson_2

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

func if_short() {
	fmt.Print("3^2 < 20 == ", funcPow(3, 2, 20))
}
