package maps_problems

func ValidAnagrams(s, t string) bool {
	runes1 := []rune(s)
	runes2 := []rune(t)

	if len(runes1) != len(runes2) {
		return false
	}

	seen := make(map[rune]int)

	for _, elem := range runes1 {
		seen[elem]++
	}

	for _, elem := range runes2 {
		seen[elem]--
		if seen[elem] < 0 {
			return false
		}
	}

	return true

}

// TODO: dodelat' blyat'
