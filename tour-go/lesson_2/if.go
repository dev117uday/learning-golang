package lesson_2

import (
	"fmt"
	"math"
)

func funcSqrt(x float64) string {
	if x < 0 {
		return funcSqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func if_condition() {
	fmt.Print(funcSqrt(-40))
}
