package main

import (
	"alg/data_structrue/lists"
	"fmt"
)

func main() {
	dlist := lists.NewDoublyLinkedList()

	dlist.Append(2)
	dlist.Append(4)
	dlist.Append(3)
	dlist.Append(5)

	dlist.PrintForward()
	fmt.Println(dlist.GetAt(3))
	dlist.PrintForward()

	// list2 := lists.NewLinkedList()
	// list2.Print()

	// list2.Print()
	// fmt.Println(list2.Delete(5))
	// list2.Prepend(6)
	// list2.Print()
}
