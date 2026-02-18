package two_pointers

func CalculateContainer(height []int) int {

	maxArea := 0

	left := 0
	right := len(height) - 1

	for left < right {
		area := calcArea(right-left, height[left], height[right])

		if area > maxArea {
			maxArea = area
		}

		if height[right] > height[left] {
			left++
		} else {
			right--
		}

	}

	return maxArea
}

func calcArea(x, y, y2 int) int {
	return x * min(y, y2)
}

// 1,8,6,2,5,4,8,3,7
//   l             r
