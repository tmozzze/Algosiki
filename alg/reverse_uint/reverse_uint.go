package reverse_uint

func ReverseUint(num uint) uint {
	var result uint
	for num > 0 {
		ost := num % 10
		num /= 10
		result = result*10 + ost
	}
	return result
}
