package graphs

import (
	"fmt"
	"testing"
)

var namedGraph = NamedGraph{adjMap: map[string][]string{
	"A": {"D"},
	"B": {"D"},
	"C": {"A", "B"},
	"D": {"H", "G"},
	"E": {"A", "D", "F"},
	"F": {"K", "J"},
	"G": {"I"},
	"H": {"I", "J"},
	"I": {"L"},
	"J": {"L", "M"},
	"K": {"J"},
	"L": {},
	"M": {},
}}

func TestNamedGraph_DepthFS(t *testing.T) {
	fmt.Println(namedGraph.DepthFS("C"))
}

func TestNamedGraph_YuriiTopologicalSort(t *testing.T) {
	fmt.Println(namedGraph.TopologicalSort())
}

func TestNamedGraph_TopSort(t *testing.T) {
	fmt.Println(namedGraph.TopSort())
}
