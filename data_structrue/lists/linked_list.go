package lists

import "fmt"

type Node struct {
	data interface{}
	next *Node
}

type LinkedList struct {
	head *Node
	len  int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{head: nil, len: 0}
}

func (ll *LinkedList) Append(data interface{}) {
	newNode := &Node{data: data, next: nil}

	if ll.head == nil {
		ll.head = newNode
	} else {
		current := ll.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
	ll.len++
}

func (ll *LinkedList) Prepend(data interface{}) {
	newNode := &Node{data: data, next: ll.head}
	ll.head = newNode
	ll.len++
}

func (ll *LinkedList) Delete(data interface{}) bool {
	if ll.head == nil {
		return false
	}

	if ll.head.data == data {
		ll.head = ll.head.next
		ll.len--
		return true
	}

	current := ll.head
	for current.next != nil {
		if current.next.data == data {
			current.next = current.next.next
			ll.len--
			return true
		} else {
			current = current.next
		}
	}
	return false
}

func (ll *LinkedList) Find(data interface{}) (*Node, bool) {
	current := ll.head
	for current != nil {
		if current.data == data {
			return current, true
		}
		current = current.next
	}

	return nil, false
}

func (ll *LinkedList) Len() int {
	return ll.len
}

func (ll *LinkedList) IsEmpty() bool {
	return ll.head == nil
}

func (ll *LinkedList) Print() {
	if ll.head == nil {
		fmt.Println("List is empty")
		return
	}

	current := ll.head
	for current.next != nil {
		fmt.Printf("%v --> ", current.data)
		current = current.next
	}
	fmt.Printf("%v\n", current.data)
}
