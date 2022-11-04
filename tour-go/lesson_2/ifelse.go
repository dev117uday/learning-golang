package lesson_2

import (
	"fmt"
)

func if_else() {
	grade := 90
	if grade > 90 {
		fmt.Print("Over 90")
	} else if grade <= 90 && grade > 80 {
		fmt.Print("passed")
	} else {
		fmt.Print("Fail")
	}

}
