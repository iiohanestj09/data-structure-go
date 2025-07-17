package main

import "fmt"

type Graph struct {
    vertices int
    matrix   [][]int
}

func NewGraph(size int) *Graph {
    matrix := make([][]int, size)
    for i := range matrix {
        matrix[i] = make([]int, size)
    }
    return &Graph{vertices: size, matrix: matrix}
}

func (g *Graph) AddEdge(from, to int) {
    g.matrix[from][to] = 1
    g.matrix[to][from] = 1
}

func (g *Graph) Display() {
    for _, row := range g.matrix {
        fmt.Println(row)
    }
}

func main() {
    graph := NewGraph(4)

    graph.AddEdge(0, 1)
    graph.AddEdge(1, 2)
    graph.AddEdge(2, 3)
    graph.AddEdge(3, 0)
    fmt.Println("Adjacency Matrix:")
    graph.Display()
}