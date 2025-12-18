package main

import (
	"alg/concurrency"
	"fmt"
	"sync"
	"time"
)

func main() {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 20)
	ch3 := make(chan int, 20)

	ch1 <- 1
	ch2 <- 2
	ch2 <- 3
	ch3 <- 5
	close(ch1)
	close(ch2)
	close(ch3)

	mergedCh := concurrency.MergeChansWP[int](ch1, ch2, ch3)

	for elem := range mergedCh {
		fmt.Println(elem)
	}

	splitedCh := concurrency.SplitChans[int](mergedCh, 3)
	time.Sleep(1 * time.Second)

	var wg sync.WaitGroup

	// Пробегаемся по списку каналов
	for i, ch := range splitedCh {
		wg.Add(1)

		// Для каждого канала запускаем свою читающую горутину
		go func(idx int, c chan int) {
			defer wg.Done()
			fmt.Println("ch", i)
			for elem := range c {
				fmt.Print(elem)
			}
		}(i, ch) // Передаем переменные, чтобы зафиксировать их значения
	}

	// Ждем, пока все каналы будут вычитаны и закрыты
	wg.Wait()

}
