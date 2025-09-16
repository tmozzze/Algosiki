package find_index_x

func FindIndexX(arr []int, x int) int {
	left, right := 0, len(arr)-1
	result := -1

	for left <= right {
		mid := (left + right) / 2
		if arr[mid] == x {
			result = mid
			return result
		} else if arr[mid] < x {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return result
}
