package max_subarray

func MaxSubArr(array []int) int {
	if len(array) == 0 {
		return 0
	}
	maxSum := array[0]
	currentSum := array[0]

	for i := 1; i < len(array)-1; i++ {
		if currentSum+array[i] > array[i] {
			currentSum += array[i]

		} else {
			currentSum = array[i]
		}
		if currentSum > maxSum {
			maxSum = currentSum
		}
	}
	return maxSum
}
