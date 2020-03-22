package graphs

import (
	"fmt"
	"reflect"
	"testing"
)

var testGraphList = GraphList{adjList: [][]int{
	0: {1, 3},
	1: {2, 4},
	2: {4},
	3: {},
	4: {2},
	5: {},
}}

var testGraphArr = GraphArr{adjMatrix: [][]bool{
	0: {0: false, 1: true, 2: false, 3: true, 4: false, 5: false},
	1: {0: false, 1: false, 2: true, 3: false, 4: true, 5: false},
	2: {0: false, 1: false, 2: false, 3: false, 4: true, 5: false},
	3: {0: false, 1: false, 2: false, 3: false, 4: false, 5: false},
	4: {0: false, 1: false, 2: true, 3: false, 4: false, 5: false},
	5: {0: false, 1: false, 2: false, 3: false, 4: false, 5: false},
}}

var testConnectedComponents = GraphList{adjList: [][]int{
	0:  {4, 8, 13, 14},
	1:  {5},
	2:  {9, 15},
	3:  {9},
	4:  {0, 8},
	5:  {1, 16, 17},
	6:  {7, 11},
	7:  {6, 11},
	8:  {0, 4, 14},
	9:  {2, 3, 15},
	10: {15},
	11: {6, 7},
	12: {},
	13: {0, 14},
	14: {0, 8, 13},
	15: {2, 9, 10},
	16: {5},
	17: {5},
}}

var testConnectedComponentsArr = GraphArr{adjMatrix: [][]bool{
	0:  {0: false, 1: false, 2: false, 3: false, 4: true, 5: false, 6: false, 7: false, 8: true, 9: false, 10: false, 11: false, 12: false, 13: true, 14: true, 15: false, 16: false, 17: false},
	1:  {0: false, 1: false, 2: false, 3: false, 4: false, 5: true, 6: false, 7: false, 8: false, 9: false, 10: false, 11: false, 12: false, 13: false, 14: false, 15: false, 16: false, 17: false},
	2:  {0: false, 1: false, 2: false, 3: false, 4: false, 5: false, 6: false, 7: false, 8: false, 9: true, 10: false, 11: false, 12: false, 13: false, 14: false, 15: true, 16: false, 17: false},
	3:  {0: false, 1: false, 2: false, 3: false, 4: false, 5: false, 6: false, 7: false, 8: false, 9: true, 10: false, 11: false, 12: false, 13: false, 14: false, 15: false, 16: false, 17: false},
	4:  {0: true, 1: false, 2: false, 3: false, 4: false, 5: false, 6: false, 7: false, 8: true, 9: false, 10: false, 11: false, 12: false, 13: false, 14: false, 15: false, 16: false, 17: false},
	5:  {0: false, 1: true, 2: false, 3: false, 4: false, 5: false, 6: false, 7: false, 8: false, 9: false, 10: false, 11: false, 12: false, 13: false, 14: false, 15: false, 16: true, 17: true},
	6:  {0: false, 1: false, 2: false, 3: false, 4: false, 5: false, 6: false, 7: true, 8: false, 9: false, 10: false, 11: true, 12: false, 13: false, 14: false, 15: false, 16: false, 17: false},
	7:  {0: false, 1: false, 2: false, 3: false, 4: false, 5: false, 6: false, 7: true, 8: false, 9: false, 10: false, 11: true, 12: false, 13: false, 14: false, 15: false, 16: false, 17: false},
	8:  {0: true, 1: false, 2: false, 3: false, 4: true, 5: false, 6: false, 7: false, 8: false, 9: false, 10: false, 11: false, 12: false, 13: false, 14: true, 15: false, 16: false, 17: false},
	9:  {0: false, 1: false, 2: true, 3: true, 4: false, 5: false, 6: false, 7: false, 8: false, 9: false, 10: false, 11: false, 12: false, 13: false, 14: false, 15: true, 16: false, 17: false},
	10: {0: false, 1: false, 2: false, 3: false, 4: false, 5: false, 6: false, 7: false, 8: false, 9: false, 10: false, 11: false, 12: false, 13: false, 14: false, 15: true, 16: false, 17: false},
	12: {0: false, 1: false, 2: false, 3: false, 4: false, 5: false, 6: false, 7: false, 8: false, 9: false, 10: false, 11: false, 12: false, 13: false, 14: false, 15: false, 16: false, 17: false},
	13: {0: true, 1: false, 2: false, 3: false, 4: false, 5: false, 6: false, 7: false, 8: false, 9: false, 10: false, 11: false, 12: false, 13: false, 14: true, 15: false, 16: false, 17: false},
	14: {0: true, 1: false, 2: false, 3: false, 4: false, 5: false, 6: false, 7: false, 8: true, 9: false, 10: false, 11: false, 12: false, 13: true, 14: false, 15: false, 16: false, 17: false},
	15: {0: false, 1: false, 2: true, 3: false, 4: false, 5: false, 6: false, 7: false, 8: false, 9: true, 10: true, 11: false, 12: false, 13: false, 14: false, 15: false, 16: false, 17: false},
	16: {0: false, 1: false, 2: false, 3: false, 4: false, 5: true, 6: false, 7: false, 8: false, 9: false, 10: false, 11: false, 12: false, 13: false, 14: false, 15: false, 16: false, 17: false},
	17: {0: false, 1: false, 2: false, 3: false, 4: false, 5: true, 6: false, 7: false, 8: false, 9: false, 10: false, 11: false, 12: false, 13: false, 14: false, 15: false, 16: false, 17: false},
}}

func TestConnectedComponents(t *testing.T) {
	fmt.Println(testConnectedComponents.ConnectedComponents())
}

func TestDepthFirstSearchRec(t *testing.T) {
	fmt.Println(testGraphList.DepthFirstSearch(0))
}

func TestDepthFirstSearchIter(t *testing.T) {
	fmt.Println(testGraphList.DepthFirstSearchIter(0))
}

func TestBreadthFirstSearch(t *testing.T) {
	fmt.Println(testGraphList.BreadthFirstSearch(0))
}

func TestGraphArr_DepthFirstSearch(t *testing.T) {
	fmt.Println(testGraphArr.DepthFirstSearch(0))
}

func TestGraphArr_BreathFirstSearch(t *testing.T) {
	fmt.Println(testGraphArr.BreadthFirstSearch(0))
}

func TestGraphList_ConnectedComponents(t *testing.T) {
	fmt.Println(testConnectedComponentsArr.ConnectedComponents())
}

func TestConvertToArrGraph(t *testing.T) {
	var expected = testConnectedComponents.ConnectedComponents()
	var actual = ConvertToArrGraph(&testConnectedComponents).ConnectedComponents()
	fmt.Println(reflect.DeepEqual(expected, actual))
}
