package concurrency

import (
	"sync"
)

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

func MergeChansWP[T any](chans ...chan T) chan T {
	resCh := make(chan T)
	poolChs := make(chan chan T)
	wg := sync.WaitGroup{}

	go func() {
		for _, ch := range chans {
			poolChs <- ch
		}
		close(poolChs)
	}()

	// workers
	wg.Add(5)
	for range 5 {
		go func() {
			defer wg.Done()
			for ch := range poolChs {
				for elem := range ch {
					resCh <- elem
				}
			}
		}()
	}

	go func() {
		wg.Wait()
		close(resCh)
	}()

	return resCh

}
