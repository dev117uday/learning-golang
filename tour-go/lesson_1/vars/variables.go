package main

import (
	"fmt"
	"math"
)

var c, python, java bool
var i, j int = 1, 2

// global constant
const pi = 3.14

func main() {

	//constants
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", pi, "Day")

	// use := for short shand notation

	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)

	var x, y int = 3, 4
	var f2 float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f2)
	fmt.Println(x, y, z)

	// %v	the value in a default format
	// when printing structs, the plus flag (%+v) adds field names
	// %#v	a Go-syntax representation of the value
	// %T	a Go-syntax representation of the type of the value
	// %%	a literal percent sign; consumes no value
	// %q	a single-quoted character literal safely escaped with Go syntax.
	// %f	decimal point but no exponent, e.g. 123.456
	// %p	base 16 notation, with leading 0x

}

// Data Type

// int8		8-bit signed integer
// int16	16-bit signed integer
// int32	32-bit signed integer
// int64	64-bit signed integer
// uint8	8-bit unsigned integer
// uint16	16-bit unsigned integer
// uint32	32-bit unsigned integer
// uint64	64-bit unsigned integer
// int		Both int and uint contain same size, either 32 or 64 bit.
// uint		Both int and uint contain same size, either 32 or 64 bit.
// rune		It is a synonym of int32 and also represent Unicode code points.
// byte		It is a synonym of uint8.
// uintptr	It is an unsigned integer type. Its width is not defined, but its can hold all the bits of a pointer value.

// float32	32-bit IEEE 754 floating-point number
// float64	64-bit IEEE 754 floating-point number

// complex64	Complex numbers which contain float32 as a real and imaginary component.
// complex128	Complex numbers which contain float64 as a real and imaginary component.
