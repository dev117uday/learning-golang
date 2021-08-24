package main

import "fmt"

func towerOfHanoi(x int, A string,B string,C string) {
	if x==1 {
		fmt.Println("Move disk 1 from ", A , " to " , B)
		return
	}
	towerOfHanoi(x-1,A,B,C)
	fmt.Println("Move disk",x, "from ", A ," to ", C )
	towerOfHanoi(x-1,B,C,A)
}

func main(){

	n := 3
	towerOfHanoi(n,"A","B","C")

}