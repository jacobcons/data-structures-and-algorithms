package main

import "fmt"

type Pair struct {
	Key   string
	Value interface{}
}

func (p *Pair) String() string {
	return fmt.Sprintf("(%v:%v)", p.Key, p.Value)
}

type LinkedList struct {
	Value *Pair
	Next  *LinkedList
}

func NewLinkedList(value *Pair) *LinkedList {
	return &LinkedList{value, nil}
}

func (l *LinkedList) AddFirst(value *Pair) {
	clone := &LinkedList{l.Value, l.Next}
	l.Value = value
	l.Next = clone
}

func (l *LinkedList) String() string {
	if l.Next == nil {
		return l.Value.String()
	}
	return fmt.Sprintf("%v->%v", l.Value, l.Next.String())
}

type HashMap struct {
	Buckets []*LinkedList
}

func main() {
	l := NewLinkedList(&Pair{"james", 3})
	fmt.Println(l)
	l.AddFirst(&Pair{"bob", 6})
	fmt.Println(l)
	l.AddFirst(&Pair{"leo", 1})
	fmt.Println(l)
}
