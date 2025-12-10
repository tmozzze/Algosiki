package concurrency

import "sync"

func MergeChans[T any](chans ...chan T) chan T {
	resCh := make(chan T)
	wg := sync.WaitGroup{}

	for _, ch := range chans {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for elem := range ch {
				resCh <- elem
			}
		}()
	}

	go func() {
		wg.Wait()
		close(resCh)
	}()

	return resCh

}
