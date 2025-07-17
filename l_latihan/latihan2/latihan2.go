// SISTEM NAVIGASI KOTA SEDERHANA

package main

import (
	"container/heap"
	"fmt"
)

type Graph struct {
    nodes  map[string]map[string]float64
    closed map[string]map[string]bool
}

func NewGraph() *Graph {
    return &Graph{
        nodes:  make(map[string]map[string]float64),
        closed: make(map[string]map[string]bool),
    }
}

func (g *Graph) AddEdge(from, to string, weight float64, isClosed bool) {
    if _, exists := g.nodes[from]; !exists {
        g.nodes[from] = make(map[string]float64)
    }
    g.nodes[from][to] = weight

    if _, exists := g.closed[from]; !exists {
        g.closed[from] = make(map[string]bool)
    }
    g.closed[from][to] = isClosed
}

type Item struct {
    node     string
    priority float64
    index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
    return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
    return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
    pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
    item := x.(*Item)
    *pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
    old := *pq
    n := len(old)
    item := old[n-1]
    *pq = old[0 : n-1]
    return item
}

func (g *Graph) Dijkstra(start, end string) (float64, []string) {
    pq := &PriorityQueue{}
    heap.Init(pq)

    dist := make(map[string]float64)
    prev := make(map[string]string)
    for node := range g.nodes {
        dist[node] = 1<<63 - 1
    }
    dist[start] = 0
    heap.Push(pq, &Item{node: start, priority: 0})

    for pq.Len() > 0 {
        current := heap.Pop(pq).(*Item)

        if current.node == end {
            break
        }

        for neighbor, weight := range g.nodes[current.node] {
            if g.closed[current.node][neighbor] {
                continue
            }

            alt := dist[current.node] + weight
            if alt < dist[neighbor] {
                dist[neighbor] = alt
                prev[neighbor] = current.node
                heap.Push(pq, &Item{node: neighbor, priority: alt})
            }
        }
    }

    var path []string
    node := end
    for node != "" {
        path = append([]string{node}, path...)
        node = prev[node]
    }

    return dist[end], path
}

func main() {
    g := NewGraph()
    g.AddEdge("A", "B", 5, false)
    g.AddEdge("A", "C", 10, false)
    g.AddEdge("B", "C", 3, false)
    g.AddEdge("B", "D", 2, true)
    g.AddEdge("C", "D", 1, false)
    g.AddEdge("D", "A", 7, false)

    distance, path := g.Dijkstra("A", "D")
    fmt.Printf("Jarak terpendek: %.2f\n", distance)
    fmt.Println("Rute terpendek:", path)
}