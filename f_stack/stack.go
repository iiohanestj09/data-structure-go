package main

import "fmt"

type Stack struct {
    items []int
}

func NewStack() *Stack {
    return &Stack{items: []int{}}
}

func (s *Stack) Push(item int) {
    s.items = append(s.items, item)
    fmt.Printf("Elemen %d ditambahkan ke stack\n", item)
}

func (s *Stack) Pop() int {
    if len(s.items) == 0 {
        fmt.Println("Stack kosong!")
        return -1
    }
    top := s.items[len(s.items)-1]
    s.items = s.items[:len(s.items)-1]
    fmt.Printf("Elemen %d dihapus dari stack\n", top)
    return top
}

func (s *Stack) Display() {
    if len(s.items) == 0 {
        fmt.Println("Stack kosong!")
        return
    }
    fmt.Println("Isi stack (dari bawah ke atas):")
    for i, v := range s.items {
        fmt.Printf("[%d] %d\n", i, v)
    }
}

func main() {
    stack := NewStack()

    stack.Push(10)
    stack.Push(20)
    stack.Push(30)

    stack.Display()

    stack.Pop()
    stack.Display()
}
