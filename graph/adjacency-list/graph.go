package adjacency_list

type Graph struct {
	List map[int][]int
}

func NewGraph() *Graph {
	return &Graph{make(map[int][]int)}
}

func (g *Graph) AddEdge(v int, w int) {
	g.List[v] = append(g.List[v], w)
	g.List[w] = append(g.List[w], v)
}

func (g *Graph) DFS(v int) *PathData {
	pathData := NewPathData(v)
	var dfsHelper func(v int)
	dfsHelper = func(v int) {
		pathData.Visited[v] = true
		for _, w := range g.List[v] {
			_, hasVisited := pathData.Visited[w]
			if hasVisited == false {
				pathData.EdgeTo[w] = v
				dfsHelper(w)
			}
		}
	}
	dfsHelper(v)
	return pathData

}

func (g *Graph) BFS(v int) *PathData {
	pathData := NewPathData(v)
	return pathData

}

type PathData struct {
	Source  int
	Visited map[int]bool
	EdgeTo  map[int]int
}

func NewPathData(v int) *PathData {
	return &PathData{v, make(map[int]bool), make(map[int]int)}
}

func HasPathTo(pathData *PathData, v int) bool {
	return pathData.Visited[v]
}

func PathTo(pathData *PathData, v int) []int {
	if HasPathTo(pathData, v) == false {
		return nil
	}

	path := []int{}
	for v != pathData.Source {
		path = prepend(path, v)
		v = pathData.EdgeTo[v]
	}
	path = prepend(path, pathData.Source)
	return path
}

func prepend[T any](arr []T, item T) []T {
	return append([]T{item}, arr...)
}
