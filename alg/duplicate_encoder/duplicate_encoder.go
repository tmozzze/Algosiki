package duplicate_encoder

// "din"      =>  "((("
// "recede"   =>  "()()()"
// "Success"  =>  ")())())"
// "(( @"     =>  "))(("
import (
	"strings"
)

func DuplicateEncode(word string) string {
	data := make(map[rune]int)
	var newWordRunes []rune
	word = strings.ToLower(word)

	for _, char := range word {
		data[char]++
	}
	for _, char := range word {
		if data[char] != 1 {
			newWordRunes = append(newWordRunes, ')')
		} else {
			newWordRunes = append(newWordRunes, '(')
		}
	}
	return string(newWordRunes)
}
