package arrays_problems

func FindIntersectionValues(nums1 []int, nums2 []int) []int {
	set1 := make(map[int]int)
	set2 := make(map[int]int)

	for _, elem := range nums1 {
		set1[elem]++
	}

	for _, elem := range nums2 {
		set2[elem]++
	}

	count1 := 0

	for k, v1 := range set1 {
		v2 := set2[k]
		if v2 > 0 {
			count1 += v1
		}
	}

	count2 := 0

	for k, v2 := range set2 {
		v1 := set1[k]
		if v1 > 0 {
			count2 += v2
		}
	}

	return []int{count1, count2}
}
