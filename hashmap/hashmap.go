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
	hash := generateHash(key, h.capacity)
	current := h.buckets[hash]

	for current != nil {
		if current.value.Key == key {
			current.value.Value = value
			return
		}
		current = current.next
	}

	h.buckets[hash] = &linkedList{&Pair{key, value}, h.buckets[hash]}
}

func (h *HashMap) Get(key string) (interface{}, bool) {
	hash := generateHash(key, h.capacity)
	current := h.buckets[hash]

	for current != nil {
		if current.value.Key == key {
			return current.value.Value, true
		}
		current = current.next
	}

	return nil, false
}

func (h *HashMap) Remove(key string) bool {
	hash := generateHash(key, h.capacity)
	current := h.buckets[hash]
	var prev *linkedList = nil

	for current != nil {
		if current.value.Key == key {
			if current == h.buckets[hash] {
				// head of list
				h.buckets[hash] = current.next
			} else {
				// middle or end of list
				prev.next = current.next
			}
			return true
		}
		prev = current
		current = current.next
	}

	return false
}

// convert each character to ascii and write as base 31
func generateHash(key string, capacity uint32) uint32 {
	var hash uint32 = 0
	for i := 0; i < len(key); i++ {
		hash = 31*hash + uint32(key[i])
	}
	return hash % capacity
}

func populatedHashMap() *HashMap {
	h := NewHashMap()

	h.Put("bob", 3)
	h.Put("james", 2)
	h.Put("kat", 4)
	h.Put("leo", 1)
	h.Put("dylan", 1)
	h.Put("mason", 10)
	return h
}

func main() {
	// put
	h := populatedHashMap()

	// get
	fmt.Println("get tests")
	strsToGet := []string{"bob", "mason", "leo", "kat", "mike"}
	for _, str := range strsToGet {
		fmt.Println(h.Get(str))
	}

	// remove
	fmt.Println("\nremove tests")
	fmt.Println(h)
	strsToRemove := []string{"leo", "kat", "mason", "bob", "mike"}
	for _, str := range strsToRemove {
		h.Remove(str)
		fmt.Println(h)
		h = populatedHashMap()
	}
}
