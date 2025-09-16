package main

import "fmt"

func main() {
	n1 := 90
	n2 := 84
	fmt.Println(EuclideanAlg(n1, n2))
}

func EuclideanAlg(n1, n2 int) int {
	rem := 1

	if n1 > n2 {
		for rem != 0 {
			rem = n1 % n2
			n1 = n2
			n2 = rem
		}
		return n2
	}
	rem = n2 % n1
	n2 = n1
	n1 = rem

	return n1
}
