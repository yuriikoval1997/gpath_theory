package graphs

type NamedGraph struct {
	adjMap map[string][]string
}

func (graph *NamedGraph) DepthFS(startVertex string) []string {
	if _, ok := graph.adjMap[startVertex]; !ok {
		return []string{}
	}
	var visits = make(map[string]bool)
	for k := range graph.adjMap {
		visits[k] = false
	}
	var visitedVertex = make([]string, 0, len(graph.adjMap))
	var dfs func(vertex string)
	dfs = func(vertex string) {
		if visits[vertex] {
			return
		}
		visits[vertex] = true
		visitedVertex = append(visitedVertex, vertex)
		for _, v := range graph.adjMap[vertex] {
			dfs(v)
		}
	}
	dfs(startVertex)
	return visitedVertex
}

func (graph *NamedGraph) TopologicalSort() []string {
	var visits = make(map[string]bool)
	for k := range graph.adjMap {
		visits[k] = false
	}
	var topSorted = make([]string, len(graph.adjMap))
	var countdown = len(graph.adjMap) - 1
	var dfs func(vertex string)
	dfs = func(vertex string) {
		if visits[vertex] {
			return
		}
		visits[vertex] = true
		for _, v := range graph.adjMap[vertex] {
			dfs(v)
		}
		topSorted[countdown] = vertex
		countdown--
	}
	for k := range graph.adjMap {
		dfs(k)
	}
	return topSorted
}

// You can also do checks before calling the dfs(), it's cheaper than opening a stack
func (graph *NamedGraph) TopSort() []string {
	var visits = make(map[string]bool)
	for k := range graph.adjMap {
		visits[k] = false
	}
	var topSorted = make([]string, len(graph.adjMap))
	var countdown = len(graph.adjMap) - 1
	var dfs func(vertex string)
	dfs = func(vertex string) {
		visits[vertex] = true
		for _, v := range graph.adjMap[vertex] {
			if !visits[v] {
				dfs(v)
			}
		}
		topSorted[countdown] = vertex
		countdown--
	}
	for k := range graph.adjMap {
		if !visits[k] {
			dfs(k)
		}
	}
	return topSorted
}
