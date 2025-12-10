package stack_problems

func ClimbStairs(n int) int {
	if n < 2 {
		return n
	}
	a := 1
	b := 2

	for i := 3; i <= n; i++ {
		b, a = a+b, b
	}
	return b
}
