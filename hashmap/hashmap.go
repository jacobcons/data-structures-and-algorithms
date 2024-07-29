package main

import (
	"fmt"
	"strings"
)

type Pair struct {
	Key   string
	Value interface{}
}

func (p *Pair) String() string {
	return fmt.Sprintf("(%v:%v)", p.Key, p.Value)
}

type linkedList struct {
	value *Pair
	next  *linkedList
}

func (l *linkedList) String() string {
	if l.next == nil {
		return l.value.String()
	}
	return fmt.Sprintf("%v->%v", l.value, l.next.String())
}

type HashMap struct {
	buckets  []*linkedList
	capacity uint32
}

func NewHashMap() *HashMap {
	const INITIAL_CAPACITY = 4
	return &HashMap{make([]*linkedList, INITIAL_CAPACITY), INITIAL_CAPACITY}
}

func (h *HashMap) String() string {
	var sb strings.Builder
	for i := uint32(0); i < h.capacity; i++ {
		sb.WriteString(fmt.Sprintf("%v: %v\n", i, h.buckets[i]))
	}
	return sb.String()
}

func (h *HashMap) Put(key string, value interface{}) {
	index := generateHashCode(key) % h.capacity

	current := h.buckets[index]
	for current != nil {
		if current.value.Key == key {
			current.value.Value = value
			return
		}
		current = current.next
	}

	h.buckets[index] = &linkedList{&Pair{key, value}, h.buckets[index]}
}

// convert each character to ascii and write as base 31
func generateHashCode(key string) uint32 {
	var hash uint32 = 0
	for i := 0; i < len(key); i++ {
		hash = 31*hash + uint32(key[i])
	}
	return hash
}

func main() {
	h := NewHashMap()
	fmt.Println(h)

	h.Put("bob", 3)
	fmt.Println(h)

	h.Put("james", 2)
	fmt.Println(h)

	h.Put("kat", 4)
	fmt.Println(h)

	h.Put("leo", 1)
	fmt.Println(h)
}
