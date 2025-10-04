package lists

import "fmt"

type DNode struct {
	data interface{}
	prev *DNode
	next *DNode
}

type DoublyLinkedList struct {
	head *DNode
	tail *DNode
	len  int
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{head: nil, tail: nil, len: 0}
}

func (dll *DoublyLinkedList) Append(data interface{}) {
	newNode := &DNode{data: data, prev: nil, next: nil}

	if dll.head == nil {
		dll.head = newNode
		dll.tail = newNode
	} else {
		newNode.prev = dll.tail
		dll.tail.next = newNode
		dll.tail = newNode
	}
	dll.len++
}

func (dll *DoublyLinkedList) Prepend(data interface{}) {
	newNode := &DNode{data: data, prev: nil, next: nil}

	if dll.head == nil {
		dll.head = newNode
		dll.tail = newNode
	} else {
		newNode.next = dll.head
		dll.head.prev = newNode
		dll.head = newNode
	}
	dll.len++
}

func (dll *DoublyLinkedList) Delete(data interface{}) bool {
	if dll.head == nil {
		return false
	}
	if (dll.len == 1) && (dll.head.data == data) {
		dll.head = nil
		dll.tail = nil
		dll.len--
		return true
	}

	current := dll.head
	for current != nil {
		if current.data == data {
			if current == dll.head {
				dll.head = current.next
				if dll.head != nil {
					dll.head.prev = nil
				}
			} else if current == dll.tail {
				dll.tail = current.prev
				if dll.tail != nil {
					dll.tail.next = nil
				}
			} else {
				current.prev.next = current.next
				current.next.prev = current.prev
			}
			dll.len--
			return true
		}
		current = current.next
	}
	return false
}

func (dll *DoublyLinkedList) PrintForward() {
	if dll.head == nil {
		fmt.Println("List is empty")
		return
	}
	fmt.Print("Forward: ")

	current := dll.head

	for current.next != nil {
		fmt.Printf("%v --> ", current.data)
		current = current.next
	}
	fmt.Printf("%v\n", current.data)
}

func (dll *DoublyLinkedList) PrintBackward() {
	if dll.tail == nil {
		fmt.Println("List is empty")
		return
	}
	fmt.Print("Backward: ")

	current := dll.tail

	for current.prev != nil {
		fmt.Printf("%v --> ", current.data)
		current = current.prev
	}
	fmt.Printf("%v\n", current.data)

}
