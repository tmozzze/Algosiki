package hashing_problems

const (
	ZERO = "0"
	ONE  = "1"
)

func CheckOnesSegment(s string) bool {
	flag := false // finded zero

	for i, r := range s {
		sr := string(r)

		if i == 0 && sr != ONE {
			return false
		}

		if sr == ZERO {
			flag = true
		}

		if !flag {
			continue
		}

		if flag && sr == ONE {
			return false
		}

	}

	return true
}
