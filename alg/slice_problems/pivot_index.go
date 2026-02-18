package slice_problems

func PivotIndex(nums []int) int {
	totalSum := 0
	leftSum := 0

	for _, elem := range nums {
		totalSum += elem
	}

	for i, elem := range nums {

		rightSum := totalSum - leftSum - elem

		if rightSum == leftSum {
			return i
		}

		leftSum += elem

	}

	return -1

}

// 1, 7, 3, 6, 5, 6
//       l
//             r
//          mid
