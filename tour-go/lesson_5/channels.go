package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	
	s := []int{7,2,8,9,4,0,2,3,3,3,5,7}	
	c := make(chan int,2)

	// making a channel of return type int
	go sum(s[len(s)/2:], c)
	go sum(s[:len(s)/2], c)	

	y,x := <-c, <-c // receive from c	
	fmt.Println(x, y, x+y)
	
}
