package flow

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

type Graph struct {
	Id   int
	Name string
	Node []int
	Edge []int
}

type emptyInterface struct {
	typ  *struct{}
	word unsafe.Pointer
}

var g Graph
var t = reflect.TypeOf(g)

func AutoTrans() Graph {
	var input interface{}
	input = Graph{
		Id:   0,
		Name: "graph",
		Node: []int{0, 1, 2, 3, 4},
		Edge: []int{0, 1, 2, 3, 4},
	}

	graph := Graph{
		Id:   input.(Graph).Id,
		Name: input.(Graph).Name,
		Node: input.(Graph).Node,
		Edge: input.(Graph).Edge,
	}
	//graph = input	// 自动转换
	return graph
	//fmt.Println(input)
	//fmt.Println(graph)
}

func AutoAssign() interface{} {
	var input interface{}
	g = Graph{
		Id:   0,
		Name: "graph",
		Node: []int{0, 1, 2, 3, 4},
		Edge: []int{0, 1, 2, 3, 4},
	}
	input = g
	return input
}

func NewAssign() interface{} {
	var input interface{}
	g = Graph{
		Id:   0,
		Name: "graph",
		Node: []int{0, 1, 2, 3, 4},
		Edge: []int{0, 1, 2, 3, 4},
	}
	input = *(*Graph)(unsafe.Pointer(&g))
	return input
}

func NewTrans() Graph {
	var input interface{}
	input = Graph{
		Id:   0,
		Name: "graph",
		Node: []int{0, 1, 2, 3, 4},
		Edge: []int{0, 1, 2, 3, 4},
	}

	ptr := uintptr((*emptyInterface)(unsafe.Pointer(&input)).word)
	graph := *(*Graph)(unsafe.Pointer(ptr))

	fmt.Println(graph)

	return graph
}

func BenchmarkAuto(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		AutoAssign()
	}
}

func BenchmarkNewTrans(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		NewAssign()
	}
}

func TestNewTrans(t *testing.T) {
	NewTrans()
}
