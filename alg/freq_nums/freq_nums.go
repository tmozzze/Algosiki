package freq_nums

import "sort"

type Pair struct {
	Val  int
	Freq int
}

func SortByFreq(arr []int) []int {
	freqMap := make(map[int]int)
	for _, num := range arr {
		freqMap[num]++
	}

	pairs := make([]Pair, 0, len(freqMap))
	for val, freq := range freqMap {
		pairs = append(pairs, Pair{Val: val, Freq: freq})
	}

	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].Freq == pairs[j].Freq {
			return pairs[i].Val < pairs[j].Val
		}
		return pairs[i].Freq > pairs[j].Freq
	})

	result := make([]int, 0, len(arr))
	for _, p := range pairs {
		for i := 0; i < p.Freq; i++ {
			result = append(result, p.Val)
		}
	}
	return result
}
