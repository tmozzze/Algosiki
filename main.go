package main

import (
	"alg/alg/two_pointers"
)

func main() {
	var a_array = []int{1, 2, 10}
	var b_array = []int{8, 20, 30}

	//sliding_window.SlidingWindow(a_array, 3)
	//sliding_window.SliceSlidingWindow(b_array, 2)
	two_pointers.TwoPointers(a_array, b_array)

}
