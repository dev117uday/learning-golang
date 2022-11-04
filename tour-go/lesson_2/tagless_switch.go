//Switch without a condition is the same as switch true.
//This construct can be a clean way to write long if-then-else chains.

package lesson_2

import (
	"fmt"
	"time"
)

func tagless_switch() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
