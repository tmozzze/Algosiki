package slice_problems

import "fmt"

func MoveZeros(nums []int) []int {
	left := 0

	for right := 0; right < len(nums); right++ {
		fmt.Println(nums)
		if nums[right] != 0 {
			// Go позволяет менять переменные местами в одну строку
			// Если left == right, это просто обмен самого с собой (безопасно)
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
	}

	return nums
}

// 0, 1, 0, 3, 12
// 1, 3, 12, 0, 0
//           l
//              r
