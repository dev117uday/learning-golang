package lesson_1

import "fmt"

var c, python, java bool
var i, j int = 1, 2

func variables() {
	var i int
	fmt.Println(i, c, python, java)

	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)

	// use := for short shand notation
}
