package arrays_problems

import (
	"fmt"
	"sort"
)

func IntersectionOfMultipleArr(nums [][]int) {
	seen := make(map[int]int)
	var result []int

	for _, arr := range nums {
		for _, elem := range arr {
			seen[elem]++
			if seen[elem] == len(nums) {
				result = append(result, elem)
			}
		}
	}

	sort.Ints(result)

	fmt.Println(result)
}

// 3,1,2,4,5
// 1,2,3,4
// 3,4,5,6
