package lesson_2

import (
	"fmt"
	"time"
)

func switch2() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	fmt.Print(time.Saturday, "\n", today, "\n")
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}
