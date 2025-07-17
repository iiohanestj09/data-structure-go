package main

import "fmt"

type Node struct {
    Value    string
    Children []*Node
}

func (n *Node) AddChild(child *Node) {
    n.Children = append(n.Children, child)
}

func PrintTree(n *Node, level int) {
    if n == nil {
        return
    }

    for i := 0; i < level; i++ {
        fmt.Print(" ")
    }
    fmt.Println(n.Value)

    for _, child := range n.Children {
        PrintTree(child, level+1)
    }
}

func main() {
    root := &Node{Value: "A"}

    b := &Node{Value: "B"}
    c := &Node{Value: "C"}
    d := &Node{Value: "D"}

    root.AddChild(b)
    root.AddChild(c)
    root.AddChild(d)

    e := &Node{Value: "E"}
    f := &Node{Value: "F"}
    b.AddChild(e)
    b.AddChild(f)

    g := &Node{Value: "G"}
    c.AddChild(g)

    h := &Node{Value: "H"}
    i := &Node{Value: "I"}
    j := &Node{Value: "J"}
    d.AddChild(h)
    d.AddChild(i)
    d.AddChild(j)

    fmt.Println("Struktur Tree:")
    PrintTree(root, 0)
}