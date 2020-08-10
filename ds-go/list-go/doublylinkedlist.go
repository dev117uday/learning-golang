package main

import "fmt"

type DoubleLinkedList struct {
	number int
	next *DoubleLinkedList
	prev *DoubleLinkedList
}

var head *DoubleLinkedList = nil
var guide * DoubleLinkedList = nil

func insert(number int) {
	var ptr = DoubleLinkedList{}
	ptr.number = number
	ptr.next = nil
	if head == nil {
		head = &ptr
		guide = &ptr
		ptr.prev = nil
	} else {
		guide.next = &ptr
		ptr.prev = guide
		guide = &ptr
	}
}

func display(){

	var ptr *DoubleLinkedList = head
	 for ptr != nil {
	 	fmt.Println(ptr.number)
	 	ptr = ptr.next
	 }

}

func display_reverse() {

		var ptr *DoubleLinkedList = guide
		for ptr!=nil{
			fmt.Println(ptr.number)
			ptr = ptr.prev
		}

}

func main() {
	insert(0)
	insert(1)
	insert(2)
	insert(3)
	display()
	fmt.Println("")
	display_reverse()
}