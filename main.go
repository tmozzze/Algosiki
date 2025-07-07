package main

import (
	"github.com/tmozzze/Algosiki/tree/main/alg/sliding_window"
)

func main() {
	var array = []int{1, 2, 3, 4, 5, 6, 7}

	sliding_window.SlidingWindow(array, 3)
	sliding_window.SliceSlidingWindow(array, 5)

}
