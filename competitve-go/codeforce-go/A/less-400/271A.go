package main

import (
	"fmt"
)

func main() {
	var year uint16
	_, _ = fmt.Scanln(&year)
	for {
		year++
		one := uint16(year / 1000)
		two := uint16((year%1000 - year%100) / 100)
		three := uint16((year/10 - one*100) % 10)
		four := uint16(year % 10)

		if one == two || one == three || one == four || two == three || two == four || three == four {
			continue
		} else {
			fmt.Print(year)
			break
		}
	}
}
