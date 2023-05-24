package main

import (
	"fmt"
	"strings"
)

// An array has a fixed size. A slice, on the other hand, is a dynamically-sized,
// flexible view into the elements of an array. In practice,
// slices are much more common than array

func main() {

	a := make([]int, 5)
	printSlice(a)

	b := make([]int, 0, 5)
	printSlice(b)

	c := b[:2]
	printSlice(c)

	d := c[2:5]
	printSlice(d)

	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// magic statement
	s = s[:0]
	printSlice(s)

	// trim the end permanently
	s = s[:4]
	printSlice(s)

	// trim the beginning
	s = s[2:]
	printSlice(s)

	s = s[0:]
	printSlice(s)

	var s1 []int
	fmt.Println(s1, len(s1), cap(s1))
	if s1 == nil {
		fmt.Println("nil!")
	}

	s2 := []struct {
		i int
		b bool
	}{{2, true}, {3, false}, {5, true}, {7, true}, {11, false}, {13, true}}
	fmt.Println(s2)

	var s3 []int
	s3 = append(s3, 2, 3, 4)
	printSlice(s3)

}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %#v\n", len(s), cap(s), s)
}

func sliceofslice() {
	// Create a tic-tac-toe board.
	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}
