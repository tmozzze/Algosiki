package main

import (
	"alg/data_structrue/lists"
	"fmt"
)

func main() {
	list := lists.NewLinkedList()

	list.Append(2)
	list.Append(4)
	list.Append(3)
	list.Append(5)

	list2 := lists.NewLinkedList()
	list2.Print()

	list.Print()
	fmt.Println(list.Delete(4))
	list.Prepend(6)
	list.Print()
}
