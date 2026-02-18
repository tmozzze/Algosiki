package ozon_problems

func CheckMonotone(nums []int) bool {
	isMonotone := false
	isNotMonotone := false

	for i, elem := range nums {
		if i > 0 {
			if elem > nums[i-1] {
				isMonotone = true
			}
			if elem < nums[i-1] {
				isNotMonotone = true
			}
		}
	}

	if isNotMonotone && isMonotone {
		return false
	} else {
		return true
	}
}
