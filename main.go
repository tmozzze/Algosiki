package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//find_phrase.FindPhrase("Для")
	//fmt.Println(duplicate_encoder.DuplicateEncode("din"))
	//fmt.Println(count_ips.IPsBetween("20.0.0.10.0", "20.0.0.11.10")) //print 266

	// ch1, ch2, ch3 := merge_channels.MakeThreeChannelsWithValues()
	// merge_channels.PrintValuesFromChannel(merge_channels.MergeIntChannels(ch1, ch2, ch3))

	var wg sync.WaitGroup

	semaphore := make(chan struct{}, 3)

	wg.Add(10)

	for i := 0; i <= 10; i++ {
		semaphore <- struct{}{}
		go printNumber(i, &wg)
		<-semaphore
	}

	wg.Wait()

}

func printNumber(value int, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Second * 1)
	fmt.Println(value)
}
