package main

import "fmt"

func main() {

	var number int = 0
	_, _ = fmt.Scanln(&number)
	 var matrix = make([][2]int,number)

	 var max,min = 0,0

	 for i:=0; i<number; i++ {
		 _, _ = fmt.Scan(&matrix[i][0])
		 _, _ = fmt.Scan(&matrix[i][1])
		 min =  matrix[i][1] - matrix[i][0] + min
		if min > max {
			max = min
		}
	 }
	 fmt.Println(max)
}
