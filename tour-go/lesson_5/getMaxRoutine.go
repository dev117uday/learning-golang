package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Test goroutine")
	fmt.Println(runtime.GOMAXPROCS(-1))
}
