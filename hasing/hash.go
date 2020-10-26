package main

import (
	"fmt"
	"sort"
)

var (
	minLoadFactor    = 0.25
	maxLoadFactor    = 0.75
	defaultTableSize = 3
)

type Element struct {
	key   int
	value int
	next  *Element
}

type Hash struct {
	buckets []*Element
}

type HashTable struct {
	table *Hash
	size  *int
}

func createHashTable(tableSize int) HashTable {
	num := 0
	hash := Hash{make([]*Element, tableSize)}
	return HashTable{table: &hash, size: &num}
}

func CreateHashTable() HashTable {
	num := 0
	hash := Hash{make([]*Element, defaultTableSize)}
	return HashTable{table: &hash, size: &num}
}

func hashFuntion(key int, size int) int {
	return key % size
}

func (h *HashTable) Display() {
	fmt.Printf("-------- %d elements ------------\n", *h.size)
	for i, node := range h.table.buckets {
		fmt.Printf("%d: ", i)
		for node != nil {
			fmt.Printf("[%d, %d]", node.key, node.value)
			node = node.next
		}
		fmt.Println("nil")
	}
}

func (h *HashTable) put(key, value int) bool {
	index := hashFuntion(key, len(h.table.buckets))
	iterator := h.table.buckets[index]
	node := Element{key, value, nil}
	if iterator == nil {
		h.table.buckets[index] = &node
	} else {
		prev := &Element{0, 0, nil}
		for iterator != nil {
			if iterator.key == key { // collision
				iterator.value = value
				return false
			}
			prev = iterator
			iterator = iterator.next
		}
		prev.next = &node
	}
	*h.size++
	return true
}

func (h *HashTable) Put(key, value int) {
	sizeChanged := h.put(key, value)
	if sizeChanged == true {
		h.checkLoadFactorAndUpdate()
	}
}

func (h *HashTable) Get(key, value int) (bool, int) {
	index := hashFuntion(key, len(h.table.buckets))
	iterator := h.table.buckets[index]
	for iterator != nil {
		if iterator.key == key {
			return true, iterator.value
		}
		iterator = iterator.next
	}
	return false, 0
}

func (h *HashTable) del(key int) bool {
	index := hashFuntion(key, len(h.table.buckets))
	iterator := h.table.buckets[index]
	if iterator == nil {
		return false
	}
	if iterator.key == key {
		h.table.buckets[index] = iterator.next
		*h.size--
		return true
	} else {
		prev := iterator
		iterator = iterator.next
		for iterator != nil {
			if iterator.key == key {
				prev.next = iterator.next
				*h.size--
				return true
			}
			prev = iterator
			iterator = iterator.next
		}
		return false
	}
}

func (h *HashTable) Del(key int) bool {
	sizeChanged := h.del(key)
	if sizeChanged == true {
		h.checkLoadFactorAndUpdate()
	}
	return sizeChanged
}

func (h *HashTable) getLoadFactor() float64 {
	return float64(*h.size) / float64(len(h.table.buckets))
}

func (h *HashTable) checkLoadFactorAndUpdate() {
	if *h.size == 0 {
		return
	} else {
		loadFactor := h.getLoadFactor()
		if loadFactor < minLoadFactor {
			fmt.Println("** Loadfactor below limit, reducing hashtable size **")
			hash := createHashTable(len(h.table.buckets) / 2)
			for _, record := range h.table.buckets {
				for record != nil {
					hash.put(record.key, record.value)
					record = record.next
				}
			}
			h.table = hash.table
		} else if loadFactor > maxLoadFactor {
			fmt.Println("** Loadfactor above limit, increasing hashtable size **")
			hash := createHashTable(*h.size * 2)
			for _, record := range h.table.buckets {
				for record != nil {
					hash.put(record.key, record.value)
					record = record.next
				}
				h.table = hash.table
			}
		}
	}
}

type sortRune []rune

func (s sortRune) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRune) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRune) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRune(r))
	return string(r)
}

func removeChars(str, charsToBeRemoved string) string {
	mymap := map[byte]bool{}
	result := ""
	for i := 0; i < len(charsToBeRemoved); i++ {
		mymap[charsToBeRemoved[i]] = true
	}

	for i := 0; i < len(str); i++ {
		if mymap[str[i]] == false {
			result += string(str[i])
		}
	}
	return result
}

func firstUniqChar(str string) rune {
	for i := 0; i < len(str); i++ {
		repeated := false
		for j := 0; j < len(str); j++ {
			if i != j && str[i] == str[j] {
				repeated = true
				break
			}
		}
		if !repeated {
			return rune(str[i])
		}
	}
	return rune(0)
}

func firstUniqCharHT(str string) rune {
	m := make(map[rune]uint, len(str))
	for _, r := range str {
		m[r]++
	}
	for _, r := range str {
		if m[r] == 1 {
			return r
		}
	}
	return rune(0)
}

func firstRepeatedChar(str string) byte {
	n := len(str)
	counters := [256]int{}

	//init
	for i := 0; i < 256; i++ {
		counters[i] = 0
	}

	for i := 0; i < n; i++ {
		if counters[str[i]] == 1 {
			return str[i]
		}
		counters[str[i]]++
	}

	return byte(0)

}

func main() {

	h := CreateHashTable()
	h.Display()

	h.Put(1, 2)
	h.Display()

	h.Put(2, 3)
	h.Display()

	h.put(3, 4)
	h.Display() // full

	h.Put(4, 5)
	h.Display() // increasing

	h.Put(5, 6)
	h.Display()

	h.Del(1)
	h.Display()

	h.Del(2)
	h.Display()

	h.Del(3)
	h.Display()

	h.Put(3, 4)
	h.Display()

	h.Put(4, 5)
	h.Display()

	h.Put(5, 6)
	h.Display()

	h.Del(4)
	h.Display()

	h.Del(5)
	h.Display() // reducing

	h.Put(11, 12)
	h.Display() // add index 3

	h.Put(12, 13)
	h.Display() // add index 0

	h.Put(13, 14) //increasing
	h.Display()

	h.Put(14, 15)
	h.Display()

	h.Put(15, 16)
	h.Display()

}
