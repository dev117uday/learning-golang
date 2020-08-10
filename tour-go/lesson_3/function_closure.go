// Go functions may be closures. A closure is a function value that references variables from
// outside its body. The function may access and assign to the referenced variables; in this
// sense the function is "bound" to the variables. hence the function becomes static in some sense
// and its lifetime is that of the variable it gets binded to. also the internal variable of the func also
// lives as long as the binded variable

package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
