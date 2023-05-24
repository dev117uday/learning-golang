package main

import "fmt"

type Point struct {
	x int
	y int
}

func main() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j

	v := Point{1, 2}
	p1 := &v
	// (*p) isnt needed here, go is smart
	p1.x = 10
	fmt.Println(p1.x)

}
