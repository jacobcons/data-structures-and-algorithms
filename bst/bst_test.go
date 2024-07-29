package bst

import (
	"slices"
	"testing"
)

func PopulatedBst() *Bst[int] {
	t := NewBst(3)
	t.Insert(2)
	t.Insert(5)
	t.Insert(1)
	t.Insert(4)
	t.Insert(6)
	return t
}

func TestInsert(t *testing.T) {
	got := PopulatedBst().PreOrder()
	want := []int{3, 2, 1, 5, 4, 6}
	if slices.Equal(got, want) == false {
		t.Errorf("got=%v, want=%v", got, want)
	}
}

func TestSearch(t *testing.T) {
	tests := []struct {
		valueToSearch int
		want          bool
	}{
		{2, true},
		{6, true},
		{0, false},
		{7, false},
	}
	tree := PopulatedBst()
	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			got := tree.Search(test.valueToSearch)
			if got != test.want {
				t.Errorf("tree.Search(%v) = %v, want %v", test.valueToSearch, got, test.want)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	tests := []struct {
		valueToDelete int
		want          []int
	}{
		{12, []int{3, 2, 1, 5, 4, 6}},
		{1, []int{3, 2, 5, 4, 6}},
		{6, []int{3, 2, 1, 5, 4}},
		{2, []int{3, 1, 5, 4, 6}},
		{3, []int{4, 2, 1, 5, 6}},
	}
	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			tree := PopulatedBst()
			got := tree.Delete(test.valueToDelete).PreOrder()
			if slices.Equal(got, test.want) == false {
				t.Errorf("tree.Delete(%v) = %v, want %v", test.valueToDelete, got, test.want)
			}
		})
	}
}
