package two_pointers

import (
	"fmt"
	"unicode"
)

func ValidPalindrome(s string) bool {
	runes := []rune(s)

	left := 0
	right := len(runes) - 1

	for left < right {
		fmt.Println(left)
		if !(unicode.IsLetter(runes[left]) || unicode.IsDigit(runes[left])) {
			left++
			continue
		}
		if !(unicode.IsLetter(runes[right]) || unicode.IsDigit(runes[right])) {
			right--
			continue
		}

		if !(unicode.ToLower(runes[left]) == unicode.ToLower(runes[right])) {
			return false
		}
		left++
		right--

	}

	return true
}

// sas os as
//     lr
