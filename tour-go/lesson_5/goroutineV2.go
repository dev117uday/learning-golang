package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Testing go routine")

	go func() {
		fmt.Println("Printing inside go routine")
	}()

	time.Sleep(100 * time.Millisecond)

}
