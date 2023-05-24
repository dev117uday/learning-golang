package main

import (
	"fmt"
)

type Vertex struct {
	X, Y float64
}

type rectangle struct {
	height, width float64
}

type geometry interface {
	area() float64
	perimeter() float64
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (r rectangle) area() float64 {
	return r.height * r.width
}

func (r rectangle) perimeter() float64 {
	return 2*r.height + 2*r.width
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimeter())
}

type I interface {
	M()
}

type MyS struct {
	S string
}

func (t *MyS) M() {
	fmt.Println(t.S)
}

func main() {
	r := rectangle{3, 4}
	measure(r)

	var j interface{}
	describe(j)

	j = 42
	describe(j)

	j = "hello"
	describe(j)

	var i I           // just like java
	i = &MyS{"Hello"} // references
	describe(i)
	i.M()

}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
