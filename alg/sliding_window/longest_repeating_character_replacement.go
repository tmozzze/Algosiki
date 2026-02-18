package sliding_window

func LongestRepeatingCharReplacement(s string, k int) int {
	if s == "" {
		return 0
	}

	if len(s) == 1 {
		return 1
	}

	left := 0
	seen := make(map[rune]int)

	maxFreq := 0
	maxLen := 0
	len := 0
	replacmentNeeded := 0

	runes := []rune(s)

	for right, elem := range runes {
		seen[elem]++

		// обновляем самый частый символ в окне
		maxFreq = max(seen[elem], maxFreq)

		len = right - left + 1
		replacmentNeeded = len - maxFreq

		if replacmentNeeded > k {
			seen[runes[left]]--
			left++
			len = right - left + 1
			replacmentNeeded = len - maxFreq
		}

		maxLen = max(len, maxLen)

	}

	return maxLen
}
