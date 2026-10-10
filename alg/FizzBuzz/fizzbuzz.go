package fizzbuzz

import (
	"strconv"
	"strings"
)

func FizzBuzz(n int) string {
	var b strings.Builder

	for elem := 0; elem <= n; elem++ {
		if elem != 0 && elem%3 == 0 && elem%5 == 0 {
			b.WriteString("FizzBuzz")
			b.WriteByte(' ')
		} else if elem != 0 && elem%3 == 0 {
			b.WriteString("Fizz")
			b.WriteByte(' ')
		} else if elem != 0 && elem%5 == 0 {
			b.WriteString("Buzz")
			b.WriteByte(' ')
		} else {
			b.WriteString(strconv.Itoa(elem))
			b.WriteByte(' ')
		}
	}

	return b.String()
}
