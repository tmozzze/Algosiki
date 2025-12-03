package hashing_problems

type MyHashMap struct {
	Check   [1000001]bool
	Storage []*KeyValueStorage
}

type KeyValueStorage struct {
	Key   int
	Value int
}

func kvsNew(k, v int) *KeyValueStorage {
	return &KeyValueStorage{Key: k, Value: v}
}

func ConstructorMap() MyHashMap {
	return MyHashMap{}
}

func (m *MyHashMap) Put(key int, value int) {
	// if map contains key
	if m.Check[key] {
		for _, kvs := range m.Storage {
			if kvs.Key == key {
				kvs.Value = value
			}
		}
	} else {
		// if map doesnt contain key
		kvs := kvsNew(key, value)
		m.Storage = append(m.Storage, kvs)
		m.Check[key] = true
	}
}

func (m *MyHashMap) Get(key int) int {
	// if map contains key
	if m.Check[key] {
		for _, kvs := range m.Storage {
			if kvs.Key == key {
				return kvs.Value
			}
		}
	}
	return -1
}

func (m *MyHashMap) Remove(key int) {
	// if map contains key
	if m.Check[key] {
		for i, kvs := range m.Storage {
			if kvs.Key == key {
				m.Storage = append(m.Storage[:i], m.Storage[i+1:]...)
				m.Check[key] = false
			}
		}
	}
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */

// const N = 1000001

// type MyHashMap struct {
// 	Map [N]int
// }

// func ConstructorMap() MyHashMap {
// 	hs := MyHashMap{}
// 	hs.Map[0] = -1
// 	for i := 1; i < N; i *= 2 {
// 		copy(hs.Map[i:], hs.Map[:i])
// 	}
// 	return hs
// }

// func (hm *MyHashMap) Put(key int, value int) {
// 	hm.Map[key] = value
// }

// func (hm *MyHashMap) Get(key int) int {
// 	return hm.Map[key]
// }

// func (hm *MyHashMap) Remove(key int) {
// 	hm.Map[key] = -1
// }

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
