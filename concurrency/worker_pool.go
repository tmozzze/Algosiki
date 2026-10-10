// Сценарий: У нас есть 100 URL-адресов (просто массив строк). Нам нужно "скачать" их (сделать time.Sleep(100*time.Millisecond) и вывести в консоль). Но одновременно мы можем качать не более 5 штук (чтобы не положить сеть).
// Напиши код этого воркер пула и пришли мне сюда. (Используй каналы и WaitGroup).
package concurrency

import (
	"fmt"
	"sync"
	"time"
)

type WorkerPool struct {
	workers  int
	mu       sync.Mutex
	wg       sync.WaitGroup
	tasks    chan string
	isClosed bool
	isDone   chan struct{}
}

func NewWP(workers int) *WorkerPool {
	wp := WorkerPool{workers: workers, tasks: make(chan string, 5), isDone: make(chan struct{})}

	return &wp
}

func (wp *WorkerPool) Run() {
	wp.wg.Add(wp.workers)

	for range wp.workers {
		go func() {
			for task := range <-wp.tasks {
				defer wp.wg.Done()
				time.Sleep(100 * time.Millisecond)
				fmt.Println(task)
			}
		}()
	}

	wp.wg.Wait()
	close(wp.isDone)

}

func (wp *WorkerPool) AddTask(task string) {
	wp.mu.Lock()
	defer wp.mu.Unlock()
	if wp.isClosed {
		return
	}
	select {
	case wp.tasks <- task:
		return
	default:
		return
	}
}

func (wp *WorkerPool) Close() {
	wp.mu.Lock()
	defer wp.mu.Unlock()

	if wp.isClosed {
		return
	}

	close(wp.tasks)
}
