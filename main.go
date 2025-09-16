package main

import (
	"alg/alg/reverse_uint"
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

	// var wg sync.WaitGroup

	// semaphore := make(chan struct{}, 3)

	// wg.Add(10)

	// for i := 0; i <= 10; i++ {
	// 	semaphore <- struct{}{}
	// 	go printNumber(i, &wg)
	// 	<-semaphore
	// }

	// wg.Wait()

	//arr := []int{12, 2, 9, 7, 58, 6, 11, 8, -1, 9, 10, 9, 12, 9}
	//arr := []int{1, 2, 3, 4, 5, 6, 2, 4, 4, 8, 1, 7, 8, 9, 10}
	// fmt.Println(count_nums.CountOftenNums(arr))
	// fmt.Println(find_pair.FindPairInNotSrtdArr(arr, 11))
	//sort.Ints(arr)
	// fmt.Print(find_pair.FindPairInSrtdArr(arr, 11))
	// palindrome := "racecar"
	// notPalindrome := "hello"
	// fmt.Println(check_palindrome.CheckPalindrome(palindrome))
	// fmt.Println(check_palindrome.CheckPalindrome(notPalindrome))
	//fmt.Println(find_index_x.FindIndexX(arr, 9))
	//fmt.Println(find_index_x.FindIndexX(arr, 7))

	fmt.Println(reverse_uint.ReverseUint(1999999999999999999))
}

func printNumber(value int, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Second * 1)
	fmt.Println(value)
}
