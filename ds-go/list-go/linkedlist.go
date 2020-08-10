package main

import "fmt"

type LinkedList struct {
	number int
	next *LinkedList
}
type List struct {
	head *LinkedList
	guide *LinkedList
}

func (L *List) insert (number int) {
	var ptr = LinkedList{}
	ptr.number = number
	ptr.next = nil
	if L.head == nil {
		L.head = &ptr
		L.guide = &ptr
	} else {
		L.guide.next = &ptr
		L.guide = &ptr
	}
}

func (L *List) display() {
	var ptr *LinkedList = L.head
	for ptr != nil {
		fmt.Println(ptr.number)
		ptr = ptr.next
	}
}

func main(){
	var numberList = List {nil,nil}
	numberList.insert(0)
	numberList.insert(1)
	numberList.insert(2)
	numberList.insert(3)
	numberList.insert(4)
	numberList.display()
}