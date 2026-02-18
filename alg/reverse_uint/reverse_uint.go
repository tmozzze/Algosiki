package reverse_uint

func ReverseUint(num uint) uint {
	var result uint

	for {
		if num == 0 {
			break
		}

		result = result*10 + num%10
		num = num / 10 // 1234
	}
	return result
}

// 12345
// 5
