package hashing_problems

const MASK int = 1 // 0..001

func HasTrailingZeros(nums []int) bool {
	count := 0
	for _, num := range nums {
		if num&MASK == 0 {
			count++
		}
	}

	return count >= 2

}
