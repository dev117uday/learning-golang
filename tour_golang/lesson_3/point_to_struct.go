package main

import (
	"fmt"
)

type Point struct {
	x int
	y int
}

func main(){

	v := Point{1,2}
	p := &v
	p.x = 10
	fmt.Println(p.x)

}