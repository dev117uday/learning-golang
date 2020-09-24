package main

import "fmt"

func foo(c chan int, somevalue int) {
	c <- somevalue * 10
}

func main() {

	fooVal := make(chan int)

	go foo(fooVal, 3)
	go foo(fooVal, 5)

	x1 := <-fooVal
	x2 := <-fooVal

	fmt.Println(x1)
	fmt.Println(x2)

}
