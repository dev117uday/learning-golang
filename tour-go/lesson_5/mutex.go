package main

import (
	"fmt"
	"sync"
)

var waitGroup sync.WaitGroup
var mutex sync.RWMutex
var count = 0

func main() {

	fmt.Println("Testing mutex")

	for i := 0; i < 10; i++ {
		waitGroup.Add(2)
		mutex.RLock()
		go Print()
		mutex.Lock()
		go IncrementCount()

	}
	waitGroup.Wait()
}

func Print() {
	defer waitGroup.Done()
	defer mutex.RUnlock()
	fmt.Println("counter : ", count)
}

func IncrementCount() {
	defer waitGroup.Done()
	defer mutex.Unlock()
	count++
}
