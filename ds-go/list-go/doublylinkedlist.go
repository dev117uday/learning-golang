package main

import "fmt"

type DoubleLinkedList struct {
	number int
	next *DoubleLinkedList
	prev *DoubleLinkedList
}

type ImplDoubleLL struct {
	head *DoubleLinkedList
	guide * DoubleLinkedList
}

func (idl *ImplDoubleLL) insert(number int) {
	var ptr = DoubleLinkedList{}
	ptr.number = number
	ptr.next = nil
	if idl.head == nil {
		idl.head = &ptr
		idl.guide = &ptr
		ptr.prev = nil
	} else {
		idl.guide.next = &ptr
		ptr.prev = idl.guide
		idl.guide = &ptr
	}
}

func (idl *ImplDoubleLL) display(){

	var ptr *DoubleLinkedList = idl.head
	for ptr != nil {
		fmt.Println(ptr.number)
		ptr = ptr.next
	}

}

func (idl *ImplDoubleLL) display_reverse() {

	var ptr *DoubleLinkedList = idl.guide
	for ptr!=nil{
		fmt.Println(ptr.number)
		ptr = ptr.prev
	}

}

func main() {
	var dll = ImplDoubleLL{nil,nil}
	dll.insert(0)
	dll.insert(1)
	dll.insert(2)
	dll.insert(3)
	dll.display()
	fmt.Println("")
	dll.display_reverse()
}