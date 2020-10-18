package main

import "fmt"

func main() {

	var arr = []int{1, 0, 2, 3, 0, 4, 5, 0}
	var lenArr = len(arr) - 1

	for i := 0; i < lenArr; i++ {
		if arr[i] == 0 {
			copy(arr[i+1:], arr[i:])
			arr[i] = 0
			i++
		}

	}

	fmt.Println(arr)

}
