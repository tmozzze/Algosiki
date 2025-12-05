package hashing_problems

import (
	"fmt"
	"strconv"
)

func DigitCount(num string) bool {
	rNum := []rune(num)
	mp := make(map[int]int)

	if len(rNum) == 1 {
		return false
	}
	convInt, err := strconv.Atoi(string(rNum[0]))
	if err != nil {
		fmt.Println("zalupa v 1")
		return false
	}
	if convInt == 0 {
		return false
	}

	for _, r := range rNum {
		convInt, err := strconv.Atoi(string(r))
		if err != nil {
			fmt.Println("zalupa v 1")
			return false
		}
		mp[convInt]++
	}

	for i, r := range rNum {
		convInt, err := strconv.Atoi(string(r))
		if err != nil {
			fmt.Println("zalupa v 2")
			return false
		}
		if v, ok := mp[i]; ok {
			if v == convInt {
				continue
			}
		}
		if convInt == 0 {
			continue
		}
		return false
	}
	return true
}

// "1 2 1 0"
//  0 1 2 3
