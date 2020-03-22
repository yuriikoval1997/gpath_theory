package graphs

import (
	"graph_theory/src/data_structures"
)

type Graph interface {
	BreadthFirstSearch(startVertex int) []int
	DepthFirstSearch(startVertex int) []int
	ConnectedComponents() map[int][]int
}

// Implementation based o adjacency list
type GraphList struct {
	adjList [][]int
}

func (graph *GraphList) DepthFirstSearch(startVertex int) []int {
	if startVertex >= len(graph.adjList) {
		return []int{}
	}
	var visits = make([]bool, len(graph.adjList))
	var visitedVertices = make([]int, 0, len(graph.adjList))
	var dfs func(vertex int)
	dfs = func(vertex int) {
		if visits[vertex] {
			return
		}
		visits[vertex] = true
		visitedVertices = append(visitedVertices, vertex)
		for _, v := range graph.adjList[vertex] {
			dfs(v)
		}
	}
	dfs(startVertex)
	return visitedVertices
}

func (graph *GraphList) DepthFirstSearchIter(startVertex int) []int {
	if startVertex >= len(graph.adjList) {
		return []int{}
	}
	var visits = make([]bool, len(graph.adjList))
	var visitedVertices = make([]int, 0, len(graph.adjList))
	var stack = new(data_structures.ArrStack)
	stack.Push(startVertex)
	for !stack.IsEmpty() {
		var vertex, _ = stack.Pop()
		if visits[vertex] {
			continue
		}
		visits[vertex] = true
		visitedVertices = append(visitedVertices, vertex)
		for _, v := range graph.adjList[vertex] {
			stack.Push(v)
		}
	}
	return visitedVertices
}

// graph - adjacency list representation
// map - where the keys are the number of component
// and values are vertices belonging to the component
func (graph GraphList) ConnectedComponents() map[int][]int {
	var visits = make([]bool, len(graph.adjList))
	var components = make([]int, len(graph.adjList))
	var count = 0
	var dfs func(vertex int)
	dfs = func(vertex int) {
		if visits[vertex] {
			return
		}
		visits[vertex] = true
		components[vertex] = count
		for _, v := range graph.adjList[vertex] {
			dfs(v)
		}
	}
	for vertex := range graph.adjList {
		if !visits[vertex] {
			dfs(vertex)
			count++
		}
	}
	var res = map[int][]int{}
	for vertex, component := range components {
		res[component] = append(res[component], vertex)
	}
	return res
}

func (graph GraphList) BreadthFirstSearch(startVertex int) []int {
	if startVertex >= len(graph.adjList) {
		return []int{}
	}
	var queue = new(data_structures.ArrQueue)
	var visits = make([]bool, len(graph.adjList))
	var visitedVertices = make([]int, 0, len(graph.adjList))
	queue.Enqueue(startVertex)
	for !queue.IsEmpty() {
		var vertex, _ = queue.Dequeue()
		if visits[vertex] {
			continue
		}
		visits[vertex] = true
		visitedVertices = append(visitedVertices, vertex)
		for _, v := range graph.adjList[vertex] {
			queue.Enqueue(v)
		}
	}
	return visitedVertices
}

// Implementation based on adjacency matrix
type GraphArr struct {
	adjMatrix [][]bool
}

func (graph *GraphArr) DepthFirstSearch(startVertex int) []int {
	if startVertex >= len(graph.adjMatrix) {
		return []int{}
	}
	var visits = make([]bool, len(graph.adjMatrix))
	var visitedVertices = make([]int, 0, len(graph.adjMatrix))
	var dfs func(vertex int)
	dfs = func(vertex int) {
		if visits[vertex] {
			return
		}
		visits[vertex] = true
		visitedVertices = append(visitedVertices, vertex)
		for v, pathExists := range graph.adjMatrix[vertex] {
			if pathExists {
				dfs(v)
			}
		}
	}
	dfs(startVertex)
	return visitedVertices
}

func (graph *GraphArr) BreadthFirstSearch(startVertex int) []int {
	if startVertex >= len(graph.adjMatrix) {
		return []int{}
	}
	var visits = make([]bool, len(graph.adjMatrix))
	var visitedVertices = make([]int, 0, len(graph.adjMatrix))
	var queue = new(data_structures.ArrQueue)
	queue.Enqueue(startVertex)
	for !queue.IsEmpty() {
		var vertex, _ = queue.Dequeue()
		if visits[vertex] {
			continue
		}
		visits[vertex] = true
		visitedVertices = append(visitedVertices, vertex)
		for v, pathExists := range graph.adjMatrix[vertex] {
			if pathExists {
				queue.Enqueue(v)
			}
		}
	}
	return visitedVertices
}

func (graph *GraphArr) ConnectedComponents() map[int][]int {
	var visits = make([]bool, len(graph.adjMatrix))
	var components = make([]int, len(graph.adjMatrix))
	var count = 0
	var dfs func(vertex int)
	dfs = func(vertex int) {
		if visits[vertex] {
			return
		}
		visits[vertex] = true
		components[vertex] = count
		for v, pathExists := range graph.adjMatrix[vertex] {
			if pathExists {
				dfs(v)
			}
		}
	}
	for v := range graph.adjMatrix {
		if !visits[v] {
			dfs(v)
			count++
		}
	}
	var res = map[int][]int{}
	for v, component := range components {
		res[component] = append(res[component], v)
	}
	return res
}

type gg struct {
	adjList map[string]map[string]int //[
	// 										{
	//											vertexName: [
	//													{name: {weight}}
	//												]
	//										}
	//		  							]
}

func ConvertToArrGraph(g *GraphList) *GraphArr {
	var adjMatrix = make([][]bool, len(g.adjList))
	for i := 0; i < len(g.adjList); i++ {
		adjMatrix[i] = make([]bool, len(g.adjList))
		for _, v := range g.adjList[i] {
			adjMatrix[i][v] = true
		}
	}
	return &GraphArr{adjMatrix:adjMatrix}
}
