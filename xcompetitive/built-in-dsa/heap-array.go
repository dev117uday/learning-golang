package main

import (
	"fmt"
)

type IntegerHeap []int

func(iHeap IntegerHeap) Len() int {
	return len(iHeap)
}

func (iHeap *IntegerHeap) Push(heapintf interface{}) {
	*iHeap = append(*iHeap, heapintf.(int))
}

func (iHeap *IntegerHeap) Pop() interface{} {
	var n int
	var x1 int
	var previous IntegerHeap = *iHeap
	n = len(previous)
	x1 = previous[n-1]
	*iHeap = previous[0 : n-1]
	return x1
}

func main(){

	var intHeap *IntegerHeap = &IntegerHeap{1,4,5}
	intHeap.Push(2)
	intHeap.Push(3)
	intHeap.Push(4)
	fmt.Printf("minimum: %d\n", (*intHeap)[0])
	qa := intHeap.Pop()
	fmt.Println("Pop : ",qa)
	lenx := intHeap.Len()
	fmt.Println("Len() : ",lenx)
	fmt.Println(intHeap)

}