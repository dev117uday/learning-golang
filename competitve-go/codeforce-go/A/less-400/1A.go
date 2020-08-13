// doesnt work for bi value, refer to cpp version

package main

import (
	"fmt"
	"math"
)

func main() {
	var width float64
	var height float64
	var flagstone float64
	_, _ = fmt.Scan(&width)
	_, _ = fmt.Scan(&height)
	_, _ = fmt.Scan(&flagstone)
	one := math.Ceil(width/flagstone)
	two := math.Ceil(height/flagstone)
	fmt.Println(one*two)
}
