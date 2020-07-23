package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	//fmt.Println(a)
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f// a MyFloat implements Abser
	fmt.Println("a: ",a.Abs())
	fmt.Println("v: ",v.Abs())
	fmt.Println("f: ",f)

	a = &v // a *Vertex implements Abser
	fmt.Println(a.Abs())
	fmt.Println(v)
	fmt.Println(f)
	fmt.Println(a)

}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
