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

func (g *Graph) Adj(v int) []int {
	return g.List[v]
}
