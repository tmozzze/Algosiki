package main

import (
	"alg/concurrency"
	"fmt"
)

func main() {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 20)

	ch1 <- 1
	ch2 <- 2
	ch2 <- 3
	close(ch1)
	close(ch2)

	ch3 := concurrency.MergeChans[int](ch1, ch2)

	for elem := range ch3 {
		fmt.Println(elem)
	}

}
