package main

import "fmt"

type LinkedList struct {
	number int
	next   *LinkedList
}
type List struct {
	head  *LinkedList
	guide *LinkedList
}

func main() {
	var numberList = List{nil, nil}
	numberList.insert(0)
	numberList.insert(1)
	numberList.insert(2)
	numberList.insert(3)
	numberList.insert(4)
	numberList.insertAfter(2, 10)
	numberList.insertAfter(4, 6)
	numberList.attach(-1)
	numberList.attach(-2)
	numberList.insertBefore(6, 7)
	numberList.display()
}

func (L *List) insert(number int) {
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

func (L *List) insertAfter(number int, toSet int) {
	var ptr *LinkedList = L.head
	for ptr != nil {
		if ptr.number == number {
			if ptr.next != nil {
				var temp = LinkedList{}
				temp.number = toSet
				temp.next = ptr.next
				ptr.next = &temp
				return
			} else {
				L.insert(toSet)
			}
		}
		ptr = ptr.next
	}
}

func (L *List) insertBefore(number int, toSet int) {
	var ptr *LinkedList = L.head
	for ptr != nil {
		if ptr.next == nil {
			fmt.Println("Not Found")
			break
		} else if ptr.next.number == number {
			var newNode = LinkedList{}
			newNode.number = toSet
			newNode.next = ptr.next
			ptr.next = &newNode
			break
		}
		ptr = ptr.next
	}
}

func (L *List) attach(number int) {
	var ptr = LinkedList{}
	ptr.number = number
	ptr.next = L.head
	L.head = &ptr
}
