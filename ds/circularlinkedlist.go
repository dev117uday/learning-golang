package main

import "fmt"

type CircularLinkedList struct {
	number int
	next *CircularLinkedList
}

var head *CircularLinkedList = nil
var guide *CircularLinkedList = nil

func insert(number int) {
	var ptr = CircularLinkedList{}
	ptr.number = number
	if head == nil {
		head = &ptr
		guide = &ptr
		ptr.next = head
	} else {
		guide.next = &ptr
		ptr.next = head
		guide = &ptr
	}
}

func display(){

	var ptr = head
	fmt.Println(ptr.number)
	ptr = ptr.next
	if ptr == nil {
		return
	}
	for ptr != head {
		fmt.Println(ptr.number)
		ptr = ptr.next
	}
}

func main() {

	insert(0)
	insert(1)
	insert(2)
	insert(3)
	display()

}