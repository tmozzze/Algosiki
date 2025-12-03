package main

import (
	"alg/LeetCode/hashing_problems"
	"fmt"
)

func main() {
	hs := hashing_problems.ConstructorMap()

	hs.Put(3, 5)

	fmt.Println(hs.Get(3))
}
