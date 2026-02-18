package maps_problems

func ContainsDuplicate(nums []int) bool {
	seen := make(map[int]struct{})

	for _, elem := range nums {
		if _, ok := seen[elem]; !ok {
			seen[elem] = struct{}{}
		} else {
			return true
		}
	}
	return false
}
