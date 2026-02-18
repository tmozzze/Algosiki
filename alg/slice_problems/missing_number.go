package slice_problems

func MissingNumber(nums []int) int {
	sum := 0
	for i := 0; i < len(nums)+1; i++ {
		sum += i
	}

	for _, elem := range nums {
		sum = sum - elem
	}

	return sum
}
