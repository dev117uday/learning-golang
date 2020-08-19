package main

import "fmt"

type CircularLinkedList struct {
	number int
	next *CircularLinkedList
}

type ImpleCircularLL struct {
	head *CircularLinkedList
	guide *CircularLinkedList
}

func (icl *ImpleCircularLL) insert(number int) {
	var ptr = CircularLinkedList{}
	ptr.number = number
	if icl.head == nil {
		icl.head = &ptr
		icl.guide = &ptr
		ptr.next = icl.head
	} else {
		icl.guide.next = &ptr
		ptr.next = icl.head
		icl.guide = &ptr
	}
}
func (icl *ImpleCircularLL) display(){
	var ptr = icl.head
	fmt.Println(ptr.number)
	ptr = ptr.next
	if ptr == nil {
		return
	}
	for ptr != icl.head {
		fmt.Println(ptr.number)
		ptr = ptr.next
	}
}

func main() {
	var cirLL = ImpleCircularLL{nil,nil}
	cirLL.insert(0)
	cirLL.insert(1)
	cirLL.insert(2)
	cirLL.insert(3)
	cirLL.display()
}