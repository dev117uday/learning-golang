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
	ptr.next = L.head
	L.head = &ptr
}

func (L *List) display(){
	var ptr *LinkedList = L.head

	for ptr != nil {
		fmt.Println(ptr.number)
		ptr = ptr.next
	}
}

func main(){
	var numberFallingList = List{nil,nil}
	numberFallingList.insert(0)
	numberFallingList.insert(1)
	numberFallingList.insert(2)
	numberFallingList.insert(3)
	numberFallingList.display()
}