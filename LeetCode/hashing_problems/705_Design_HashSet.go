package hashing_problems

import "slices"

type MyHashSet struct {
	Set []int
}

func Constructor() MyHashSet {
	return MyHashSet{Set: make([]int, 0)}
}

func (hs *MyHashSet) Add(key int) {
	if !slices.Contains(hs.Set, key) {
		hs.Set = append(hs.Set, key)
	}
}

func (hs *MyHashSet) Remove(key int) {
	for i, elem := range hs.Set {
		if elem == key {
			hs.Set = append(hs.Set[:i], hs.Set[i+1:]...)
			return
		}
	}
}

func (hs *MyHashSet) Contains(key int) bool {
	return slices.Contains(hs.Set, key)
}
