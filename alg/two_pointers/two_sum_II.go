package two_pointers

func TwoSum(nums []int, target int) (int, int) {
	left := 0
	right := len(nums) - 1

	for left < right {
		if nums[left]+nums[right] == target {
			return nums[left], nums[right]
		}

		if nums[left]+nums[right] > target {
			right--
			continue
		}

		if nums[left]+nums[right] < target {
			left++
			continue
		}
	}
	return 0, 0
}

// 0 1 3 15 25       18
//     l  r
