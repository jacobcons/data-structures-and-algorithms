package bst

import (
	"cmp"
)

type Ordered = cmp.Ordered

type Bst[T Ordered] struct {
	Value T
	Left  *Bst[T]
	Right *Bst[T]
}

func NewBst[T Ordered](value T) *Bst[T] {
	return &Bst[T]{value, nil, nil}
}

func (t *Bst[T]) Insert(value T) *Bst[T] {
	if t == nil {
		return NewBst(value)
	}
	if value < t.Value {
		t.Left = t.Left.Insert(value)
	} else if value > t.Value {
		t.Right = t.Right.Insert(value)
	}
	return t
}

func (t *Bst[T]) Search(value T) bool {
	if t == nil {
		return false
	}
	if value < t.Value {
		return t.Left.Search(value)
	} else if value > t.Value {
		return t.Right.Search(value)
	}
	return true
}

func (t *Bst[T]) Delete(value T) *Bst[T] {
	if t == nil {
		return nil
	}

	if value < t.Value {
		t.Left = t.Left.Delete(value)
	} else if value > t.Value {
		t.Right = t.Right.Delete(value)
	} else {
		if t.Left == nil && t.Right == nil {
			return nil
		} else if t.Left == nil {
			return t.Right
		} else if t.Right == nil {
			return t.Left
		} else {
			current := t.Right
			for current.Left != nil {
				current = current.Left
			}
			t.Right = t.Right.Delete(current.Value)
			t.Value = current.Value
		}
	}
	return t
}

func (t *Bst[T]) PreOrder() []T {
	if t == nil {
		return nil
	}
	result := []T{t.Value}
	result = append(result, t.Left.PreOrder()...)
	result = append(result, t.Right.PreOrder()...)
	return result
}
