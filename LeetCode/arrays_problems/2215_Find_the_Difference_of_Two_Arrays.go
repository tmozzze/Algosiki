package arrays_problems

import "fmt"

func FindDifference(nums1 []int, nums2 []int) /*[][]int*/ {
	var res1, res2 []int

	set1 := make(map[int]bool)
	set2 := make(map[int]bool)

	for _, elem := range nums1 {
		set1[elem] = true
	}

	for _, elem := range nums2 {
		set2[elem] = true
	}

	for k := range set1 {
		if !set2[k] {
			res1 = append(res1, k)
		}
	}

	for k := range set2 {
		if !set1[k] {
			res2 = append(res2, k)
		}
	}

	fmt.Println(res1, res2)

}

// 1 3
// 4 5
