package hashing_problems

func NumberOfChild(n int, k int) int {
	flag := false // if true reverse
	cn := 0
	for count := 0; count < k; count++ {
		if !flag {
			cn++
			if cn == n-1 {
				flag = true
			}
		} else {
			cn--
			if cn == 0 {
				flag = false
			}
		}
	}
	return cn
}
