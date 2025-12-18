package concurrency

import (
	"sync"
)

func SplitChans[T any](in chan T, amount int) []chan T {
	resChs := make([]chan T, amount)

	// create channels
	for i := 0; i < amount; i++ {
		resChs[i] = make(chan T)
	}

	go func() {

		wg := sync.WaitGroup{}

		sema := make(chan struct{}, amount*2)

		i := 0

		for elem := range in {
			curCh := resChs[i]
			// reload counter
			i = (i + 1) % amount

			sema <- struct{}{}
			wg.Add(1)
			go func(elem T, ch chan T) {
				defer func() { <-sema }()
				defer wg.Done()

				curCh <- elem
			}(elem, curCh)

		}

		wg.Wait()
		for _, ch := range resChs {
			close(ch)
		}

	}()

	return resChs
}
