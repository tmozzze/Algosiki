package maps_problems

func TwoSum(nums []int, target int) (int, int) {
	seen := make(map[int]int)

	for i, elem := range nums {

		seenTarget := target - elem

		idx, ok := seen[seenTarget]
		if ok {
			return idx, i
		}

		seen[elem] = i
	}

	return 0, 0
}
