package main

import (
	"fmt"
)

func main() {
	s := "dvdf"
	fmt.Println("sas", lengthOfLongestSubstring(s))
}

func lengthOfLongestSubstring(s string) int {
	sRunes := []rune(s)
	if len(sRunes) == 0 {
		return 0
	} else if len(sRunes) == 1 {
		return 1
	}
	var result int
	m := make(map[rune]bool)
	left := 0
	right := 1

	m[sRunes[left]] = true
	counter := 1

	for right <= len(sRunes)-1 {
		if !m[sRunes[right]] {
			m[sRunes[right]] = true
			counter++
			right++
			fmt.Println("r", right)
			if result < counter {
				result = counter
			}
		} else {

			for m[sRunes[right]] {
				m[sRunes[left]] = false
				left++
				counter--
				fmt.Println("l", left)
			}
		}
	}
	return result
}

// dvdf
//  |
//    |
// 2
