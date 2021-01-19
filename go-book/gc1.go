package main

import (
	"fmt"
	"runtime"
	"time"
)

type data struct {
	i, j int
}

func main() {
	var N = 40000000
	var structure []data
	for i := 0; i < N; i++ {
		value := int(i)
		structure = append(structure, data{value, value})
	}
	_ = structure[0]
	runtime.GC()
	time.Sleep(time.Second)
	fmt.Println("\n", structure[0])
	time.Sleep(time.Second)
}
