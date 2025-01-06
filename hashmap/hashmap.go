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
	size     uint32
}

func NewHashMap(capacity uint32) *HashMap {
	return &HashMap{make([]*linkedList, capacity), capacity, 0}
}

func (h *HashMap) String() string {
	var sb strings.Builder
	for i := uint32(0); i < h.capacity; i++ {
		sb.WriteString(fmt.Sprintf("%v: %v\n", i, h.buckets[i]))
	}
	return sb.String()
}

func (h *HashMap) Put(key string, value interface{}) {
	const LOAD_FACTOR = 1.5
	if float32(h.size)/float32(h.capacity) >= LOAD_FACTOR {
		h.resize()
	}

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
	h.size++
}

func (h *HashMap) resize() {
	newCapacity := h.capacity * 2
	newHashMap := NewHashMap(newCapacity)
	for _, bucket := range h.buckets {
		p := bucket
		for p != nil {
			newHashMap.Put(p.value.Key, p.value.Value)
			p = p.next
		}
	}

	h.buckets = newHashMap.buckets
	h.capacity = newHashMap.capacity
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
	h := NewHashMap(4)

	h.Put("bob", 3)
	h.Put("james", 2)
	h.Put("kat", 4)
	h.Put("leo", 1)
	h.Put("dylan", 1)
	h.Put("mason", 10)
	h.Put("mason", 12)
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

	// resize
	h.Put("blue", 10)
	h.Put("red", 10)
	h.Put("yellow", 10)
	h.Put("pink", 10)
	h.Put("black", 10)
	h.Put("orange", 10)
	h.Put("purple", 10)
	h.Put("green", 10)
	fmt.Println(h)
}
