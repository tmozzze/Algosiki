package merge_channels

import (
	"fmt"
	"sync"
)

func MergeIntChannels(chs ...chan int) chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	wg.Add(len(chs))

	output := func(c chan int) {
		defer wg.Done()

		for element := range c {
			out <- element
		}
	}

	for _, ch := range chs {
		go output(ch)
	}

	go func() {
		wg.Wait()
		defer close(out)
	}()

	return out
}

func PrintValuesFromChannel(ch chan int) {
	for value := range ch {
		fmt.Println(value)
	}
}
func MakeThreeChannelsWithValues() (chan int, chan int, chan int) {
	ch1, ch2, ch3 := make(chan int), make(chan int), make(chan int)

	go func() {
		defer close(ch1)
		for _, v := range []int{1, 2, 3} {
			ch1 <- v
		}
	}()

	go func() {
		defer close(ch2)
		for _, v := range []int{10, 20, 30} {
			ch2 <- v
		}
	}()

	go func() {
		defer close(ch3)
		for _, v := range []int{100, 200} {
			ch3 <- v
		}
	}()

	return ch1, ch2, ch3
}
