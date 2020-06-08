package main

import (
	"fmt"
	"sort"
)

func miniMaxSum( arr []int) {
	sort.Ints(arr)
	fmt.Println(arr)
	mini := 0
	max := 0
	for i:=0; i<len(arr)-1; i++ {
		mini += arr[i]
	}
	for i:=1; i<len(arr); i++ {
		max += arr[i]
	}
	fmt.Println(mini," ",max)
}

func main(){
	arrx := []int{1, 3,2, 4, 5}
	miniMaxSum(arrx)

}