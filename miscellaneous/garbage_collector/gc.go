package main

import (
	"fmt"
	"runtime"
	"time"
)

type data struct {
	i, j int
}

func printStats(mem runtime.MemStats) {
	runtime.ReadMemStats(&mem)
	fmt.Println("mem.Alloc:", mem.Alloc)
	fmt.Println("mem.TotalAlloc:", mem.TotalAlloc)
	fmt.Println("mem.HeapAlloc:", mem.HeapAlloc)
	fmt.Println("mem.NumGC:", mem.NumGC)
	fmt.Println("-----")
}

func GC2() {
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

func main() {
	var mem runtime.MemStats
	for i := 0; i < 2; i++ {
		s := make([]byte, 100000000)
		if s == nil {
			fmt.Println("Operation failed!")
		}
		printStats(mem)
	}
	time.Sleep(time.Second)

	GC2()
}
