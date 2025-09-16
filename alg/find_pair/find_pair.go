package find_pair

func FindPairInNotSrtdArr(arr []int, target int) (int, int) {
	m := make(map[int]struct{})

	for _, num := range arr {
		if _, found := m[num]; found {
			return target - num, num
		}
		m[target-num] = struct{}{}

	}
	return 0, 0
}

func FindPairInSrtdArr(arr []int, target int) (int, int) {
	left := 0
	right := len(arr) - 1

	for left < right {
		sum := arr[left] + arr[right]

		if sum == target {
			return arr[left], arr[right]
		}
		if sum < target {
			left++
			continue
		}
		if sum > target {
			right--
			continue
		}
	}
	return 0, 0
}
