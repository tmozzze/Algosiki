package main

import (
	"alg/LeetCode/hashing_problems"
	"fmt"
)

func main() {
	nums := "1001"

	sortedNums := hashing_problems.CheckOnesSegment(nums)
	fmt.Println(sortedNums)

}
