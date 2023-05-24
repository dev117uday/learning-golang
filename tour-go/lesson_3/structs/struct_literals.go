package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

func (v Vertex) Abs(f int) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex) Scale(f int) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {

	v := Vertex{1, 2}
	fmt.Print(v.X)

	fmt.Println(v1, "\n", p, "\n", v2, "\n", v3)

	v2 := Vertex{3, 4}
	v2.Scale(10) // scale changes value cuz copy by reference
	v2.Abs(10)   // abs copy by value
	fmt.Println(v2)

}
