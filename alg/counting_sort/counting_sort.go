package counting_sort

func Sort(nums []int) []int {
	var maxNum int
	result := make([]int, 0, len(nums))

	// search max elem
	for _, num := range nums {
		if num > maxNum {
			maxNum = num
		}
	}

	// fiil counting array
	countingArr := make([]int, maxNum+1)
	for _, num := range nums {
		countingArr[num]++
	}

	// sort
	// missing 0 idx
	for k, v := range countingArr {
		if v == 0 {
			continue
		} else {
			if v == 1 {
				result = append(result, k)
			} else {
				for j := 1; j <= v; j++ {
					result = append(result, k)
				}
			}
		}
	}

	return result

}
