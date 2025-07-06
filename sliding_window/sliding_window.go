package sliding_window

import (
	"fmt"
)

func SlidingWindow(array []int, k int) {
	fmt.Println("Slinding window with CICLES. Window range: ", k)

	for i := 0; i < len(array)-k+1; i++ {
		window := make([]int, 0, k)

		for j := i; j < i+k; j++ {
			window = append(window, array[j])
		}
		fmt.Println(window)
	}
}

func SliceSlidingWindow(array []int, k int) {
	fmt.Println("Slinding window with SLICES. Window range: ", k)

	for i := 0; i < len(array)-k+1; i++ {
		fmt.Println(array[i : k+i])
	}
}
