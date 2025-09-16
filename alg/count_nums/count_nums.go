package count_nums

func CountOftenNums(arr []int) (int, int) {
	data := make(map[int]int)
	var max int

	for _, num := range arr {
		data[num]++
		if data[num] > data[max] {
			max = num
		}
	}
	return max, data[max]

}
