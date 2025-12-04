package main

import (
	"alg/alg/counting_sort"
	"fmt"
)

func main() {
	nums := []int{5, 1, 4, 3, 7, 19, 12, 85, 9999999, 9}

	sortedNums := counting_sort.Sort(nums)
	fmt.Println(sortedNums)

}
