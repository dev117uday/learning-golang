package main

import (
	"fmt"
)

// A goroutine is a lightweight thread managed by the Go runtime.
// channels : sends and receives block until the other side is ready.
// This allows goroutines to synchronize without explicit locks or condition variables.

// Golang provides buffered channels, which allow you to specify a fixed length of buffer capacity
// so one can send that number of data values at once

func printer[T any](s T) {
	for i := 0; i < 2; i++ {
		// time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func say(s string) {
	printer(s)
}

func sayChannel(s string, c chan int) {
	printer(s)
	c <- 1
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {

	c := make(chan int, 1) // 1 is buffer
	// channels are used to send value and receive value

	go say("world")
	// without time.Sleep, this doesnt even get time to execute
	go sayChannel("world2", c)
	say("hello")

	x := <-c
	fmt.Println(x)

	c1 := make(chan int, 10)
	go fibonacci(cap(c1), c)
	for i := range c {
		fmt.Println(i)
	}

	//The select statement lets a goroutine wait on multiple communication operations.
	// A select blocks until one of its cases can run, then it executes that case.
	// It chooses one at random if multiple are ready.

}

// v, ok := <-ch
// ok is false if there are no more values to receive and the channel is closed.

// The loop for i := range c receives values from the channel repeatedly until it is closed.

// Note: Only the sender should close a channel, never the receiver.
// Sending on a closed channel will cause a panic.

// Another note: Channels aren't like files; you don't usually need to close them.
// Closing is only necessary when the receiver must be told there are no more values coming,
// such as to terminate a range loop.
