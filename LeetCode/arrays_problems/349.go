package arrays_problems

import "fmt"

func IntersectionOfTwoArr(nums1, nums2 []int) {
	seen := map[int]struct{}{}
	var result []int

	for _, elem1 := range nums1 {
		for _, elem2 := range nums2 {
			if elem1 == elem2 {
				seen[elem1] = struct{}{}
			}
		}
	}

	for k, _ := range seen {
		result = append(result, k)
	}

	fmt.Println(result)

}
