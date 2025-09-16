package check_palindrome

import "strings"

func CheckPalindrome(strWord string) bool {
	strWord = strings.ToLower(strWord)
	runeWord := []rune(strWord)
	left := 0
	right := len(runeWord) - 1

	for left < right {
		if runeWord[left] != runeWord[right] {
			return false
		}
		left++
		right--
	}
	return true
}

//hello
