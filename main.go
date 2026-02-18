package main

import (
	"alg/alg/slice_problems"
	"fmt"
)

func main() {
	s := []int{3, 0, 1, 2, 5}

	res1 := slice_problems.MissingNumber(s)

	fmt.Println(res1)
}
