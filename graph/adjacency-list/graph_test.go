package adjacency_list

import (
	"slices"
	"testing"
)

func TestGraphApi(t *testing.T) {
	g := NewGraph()
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(4, 5)
	if slices.Equal(g.Adj(0), []int{1, 2}) == false {
		t.Errorf("expected g.Adj(0)=[1,2]")
	}
	if slices.Equal(g.Adj(4), []int{5}) == false {
		t.Errorf("expected g.Adj(4)=[5]")
	}
	if slices.Equal(g.Adj(3), []int{}) == false {
		t.Errorf("expected g.Adj(3)=nil")
	}
}
