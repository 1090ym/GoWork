package flow

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

type Graph struct {
	Id       int
	Name     string
	Label    string
	Test1    string
	Test2    string
	Test3    string
	Test4    string
	Node     []int
	Edge     []int
	SubGraph SubGraph
}

type SubGraph struct {
	Test1 string
	Test2 string
}

type emptyInterface struct {
	typ  *struct{}
	word unsafe.Pointer
}

var g Graph
var t = reflect.TypeOf(g)

// interface{}转Graph
func AutoTrans(input interface{}) Graph {
	graph := Graph{
		Id:   input.(Graph).Id,
		Name: input.(Graph).Name,
		Node: input.(Graph).Node,
		Edge: input.(Graph).Edge,
	}
	return graph
}

// 使用unsafe转换
func NewTrans(input interface{}) Graph {
	ptr := uintptr((*emptyInterface)(unsafe.Pointer(&input)).word)
	graph := *(*Graph)(unsafe.Pointer(ptr))

	fmt.Println(graph)
	return graph
}

// Graph转interface{}
func AutoAssign() interface{} {
	var input interface{}
	g = Graph{
		Id:    0,
		Name:  "graph",
		Label: "graph",
		Test1: "graph",
		Test2: "graph",
		Test3: "graph",
		Test4: "graph",
		Node:  []int{0, 1, 2, 3, 4},
		Edge:  []int{0, 1, 2, 3, 4},
	}
	input = g
	return input
}

func NewAssign() interface{} {
	var input interface{}
	g = Graph{
		Id:    0,
		Name:  "graph",
		Label: "graph",
		Test1: "graph",
		Test2: "graph",
		Test3: "graph",
		Test4: "graph",
		Node:  []int{0, 1, 2, 3, 4},
		Edge:  []int{0, 1, 2, 3, 4},
	}
	input = *(*Graph)(unsafe.Pointer(&g))
	return input
}

func BenchmarkAuto(b *testing.B) {
	var input interface{}
	input = Graph{
		Id:    0,
		Name:  "graph",
		Label: "graph",
		Test1: "graph",
		Test2: "graph",
		Test3: "graph",
		Test4: "graph",
		Node:  []int{0, 1, 2, 3, 4},
		Edge:  []int{0, 1, 2, 3, 4},
	}
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		AutoTrans(input)
	}
}

func BenchmarkNewTrans(b *testing.B) {
	var input interface{}
	input = Graph{
		Id:    0,
		Name:  "graph",
		Label: "graph",
		Test1: "graph",
		Test2: "graph",
		Test3: "graph",
		Test4: "graph",
		Node:  []int{0, 1, 2, 3, 4},
		Edge:  []int{0, 1, 2, 3, 4},
		SubGraph: SubGraph{
			Test1: "test1",
			Test2: "test2",
		},
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		NewTrans(input)
	}
}

func TestNewTrans(t *testing.T) {
	slice := make([]int, 10)
	slice = append(slice, 123)
	fmt.Println(slice)
}
