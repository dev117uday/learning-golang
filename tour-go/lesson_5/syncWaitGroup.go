// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// func main() {

// 	var waitGroup = sync.WaitGroup{}
// 	fmt.Println("Testing go routine")

// 	waitGroup.Add(1)

// 	go func() {
// 		fmt.Println("inside go routnine")
// 		time.Sleep(100 * time.Millisecond)
// 		waitGroup.Done()
// 	}()

// 	waitGroup.Wait()

// }
