package sliding_window

func CalcLongestSubstring(s string) int {

	if s == "" {
		return 0
	}

	runes := []rune(s)
	left := 0

	seen := make(map[rune]int)

	maxLen := 1

	for right, elem := range runes {
		seen[elem]++

		for seen[elem] > 1 {
			seen[runes[left]]--
			left++
		}

		maxLen = max(maxLen, right-left+1)
	}
	return maxLen

}

// a a a b c d a d c
// l r
