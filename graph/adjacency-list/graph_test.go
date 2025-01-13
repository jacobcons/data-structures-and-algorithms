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
	if slices.Equal(g.List[0], []int{1, 2}) == false {
		t.Errorf("expected g.List[0]=[1,2]")
	}
	if slices.Equal(g.List[4], []int{5}) == false {
		t.Errorf("expected g.List[4]=[5]")
	}
	if slices.Equal(g.List[3], []int{}) == false {
		t.Errorf("expected g.List[3]=[]")
	}
}

func TestDFS(t *testing.T) {
	g := createPopulatedGraph()
	pathData := g.DFS(0)

	testHasPathTo(pathData, t)

	expectedPath := []int{0, 1, 2, 5, 4, 3}
	testPathTo(pathData, expectedPath, t)
}

func TestBFS(t *testing.T) {
	g := createPopulatedGraph()
	pathData := g.BFS(0)

	testHasPathTo(pathData, t)

	expectedPath := []int{0, 1, 4, 3}
	testPathTo(pathData, expectedPath, t)
}

func createPopulatedGraph() *Graph {
	g := NewGraph()
	g.AddEdge(0, 1)
	g.AddEdge(1, 2)
	g.AddEdge(1, 4)
	g.AddEdge(2, 5)
	g.AddEdge(3, 4)
	g.AddEdge(4, 5)
	g.AddEdge(5, 6)
	g.AddEdge(5, 8)
	g.AddEdge(6, 7)
	g.AddEdge(9, 10)
	g.AddEdge(10, 11)
	return g
}

func testHasPathTo(pathData *PathData, t *testing.T) {
	connectedVertices := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	for _, v := range connectedVertices {
		if HasPathTo(pathData, v) == false {
			t.Errorf("HasPathTo should detect path from 0 to %v", v)
		}
	}

	disconnectedVertices := []int{9, 10, 11, 12}
	for _, v := range disconnectedVertices {
		if HasPathTo(pathData, v) == true {
			t.Errorf("HasPathTo shouldn't detect path from 0 to %v", v)
		}
	}
}

func testPathTo(pathData *PathData, expectedPath []int, t *testing.T) {
	actualPath := PathTo(pathData, 3)
	if slices.Equal(actualPath, expectedPath) == false {
		t.Errorf("PathTo(0, 3)=%v, expected=%v", actualPath, expectedPath)
	}

	path := PathTo(pathData, 9)
	if path != nil {
		t.Errorf("PathTo(0, 9)=%v, expected=nil", path)
	}
}
