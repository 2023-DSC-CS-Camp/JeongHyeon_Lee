package main

import (
	"container/list"
	"fmt"
	"sort"
)

func main() {
	var node, edge, start int
	fmt.Scan(&node, &edge, &start)
	graph := NewGraph(node, edge)
	DFS(graph, graph.Vertices[start-1])
	GraphInit(graph)
	fmt.Println()
	BFS(graph, graph.Vertices[start-1])
}

// 그래프 구현
type Vertex struct {
	Number   int
	Adjacent []int
	Visit    bool
}
type Graph struct {
	Vertices []*Vertex
}

func NewGraph(node, edge int) *Graph {
	var v []*Vertex
	var a, b int
	graph := &Graph{}
	for i := 1; i <= node; i++ {
		v = append(v, &Vertex{Number: i, Visit: false})
	}
	for i := 0; i < edge; i++ {
		fmt.Scan(&a, &b)
		v[a-1].Adjacent = append(v[a-1].Adjacent, v[b-1].Number)
		v[b-1].Adjacent = append(v[b-1].Adjacent, v[a-1].Number)
	}
	for i := 0; i < node; i++ {
		sort.Slice(v[i].Adjacent, func(a, b int) bool {
			return v[i].Adjacent[a] < v[i].Adjacent[b]
		})
		graph.Vertices = append(graph.Vertices, v[i])
	}
	return graph
}
func GraphInit(g *Graph) {
	for _, vertices := range g.Vertices {
		vertices.Visit = false
	}
}

//스택 구현

type Stack struct {
	l *list.List
}

func (s *Stack) Push(v *Vertex) {
	s.l.PushBack(v)
}

func (s *Stack) Pop() *Vertex {
	top := s.l.Back()
	if top != nil {
		return s.l.Remove(top).(*Vertex)
	} else {
		//리스트가 비어있음
	}
	return nil
}

func NewStack() *Stack {
	return &Stack{list.New()}
}

// DFS 구현
func DFS(graph *Graph, start *Vertex) {
	stack := NewStack()
	stack.Push(start)
	for stack.l.Len() != 0 {
		currentnode := stack.Pop()
		if currentnode.Visit == false {
			currentnode.Visit = true
			fmt.Print(currentnode.Number, " ")
			for i := len(graph.Vertices[currentnode.Number-1].Adjacent) - 1; i >= 0; i-- {
				stack.Push(graph.Vertices[graph.Vertices[currentnode.Number-1].Adjacent[i]-1])
			}
		}
	}
}

// 큐 구현
type Queue struct {
	l *list.List
}

func (q *Queue) Push(v *Vertex) {
	q.l.PushBack(v)
}
func (q *Queue) Pop() *Vertex {
	front := q.l.Front()
	if front != nil {
		return q.l.Remove(front).(*Vertex)
	}
	return nil
}
func NewQueue() *Queue {
	return &Queue{list.New()}
}

// BFS 구현
func BFS(graph *Graph, start *Vertex) {
	queue := NewQueue()
	queue.Push(start)

	for queue.l.Len() != 0 {
		currentnode := queue.Pop()
		if currentnode.Visit == false {
			currentnode.Visit = true
			fmt.Print(currentnode.Number, " ")
			for i := 0; i <= len(graph.Vertices[currentnode.Number-1].Adjacent)-1; i++ {
				queue.Push(graph.Vertices[graph.Vertices[currentnode.Number-1].Adjacent[i]-1])
			}
		}
	}
}
