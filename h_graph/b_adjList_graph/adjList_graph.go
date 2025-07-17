package main

import "fmt"

type Graph struct {
    vertices int
    adjList  map[int][]int
}

func NewGraph(size int) *Graph {
    return &Graph{vertices: size, adjList: make(map[int][]int)}
}

func (g *Graph) AddEdge(from, to int) {
    g.adjList[from] = append(g.adjList[from], to)
    g.adjList[to] = append(g.adjList[to], from)
}

func (g *Graph) Display() {
    for key, value := range g.adjList {
        fmt.Println(key, ":", value)
    }
}

func main() {
    graph := NewGraph(4)

    graph.AddEdge(0, 1)
    graph.AddEdge(1, 2)
    graph.AddEdge(2, 3)
    graph.AddEdge(3, 0)

    fmt.Println("Adjacency List:")
    graph.Display()
}
