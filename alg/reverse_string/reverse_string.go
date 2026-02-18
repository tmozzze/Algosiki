package reverse_string

func ReverseString(s string) string {
	runes := []rune(s)
	result := make([]rune, len(runes))

	for i := len(runes) - 1; i >= 0; i-- {
		result = append(result, runes[i])
	}

	return string(result)
}
