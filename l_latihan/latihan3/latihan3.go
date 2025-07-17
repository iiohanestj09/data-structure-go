// SISTEM UNDO REDO DALAM PENGOLAHAN DATA DINAMIS

package main

import (
	"fmt"
)

type Node struct {
    data int
    next *Node
}

type LinkedList struct {
    head *Node
}

type Stack struct {
    items []int
}

func (ll *LinkedList) Add(data int) {
    newNode := &Node{data: data}
    if ll.head == nil {
        ll.head = newNode
    } else {
        current := ll.head
        for current.next != nil {
            current = current.next
        }
        current.next = newNode
    }
    fmt.Println("Menambahkan:", data)
}

func (ll *LinkedList) Remove(value int, undoStack *Stack) {
    if ll.head == nil {
        fmt.Println("Tidak ada data untuk dihapus")
        return
    }
    var prev *Node
    current := ll.head

    for current != nil && current.data != value {
        prev = current
        current = current.next
    }
    if current != nil {
        undoStack.Push(current.data)
        if prev != nil {
            prev.next = current.next
        } else {
            ll.head = current.next
        }
        fmt.Println("Menghapus:", value)
    } else {
        fmt.Println("Data tidak ditemukan")
    }
}

func (ll *LinkedList) Undo(undoStack *Stack, redoStack *Stack) {
    data, success := undoStack.Pop()
    if success {
        ll.Add(data)
        redoStack.Push(data)
        fmt.Println("Undo:", data)
    }
}

func (ll *LinkedList) Redo(redoStack *Stack) {
    data, success := redoStack.Pop()
    if success {
        ll.Remove(data, &Stack{})
        fmt.Println("Redo:", data)
    }
}

func (ll *LinkedList) Print() {
    current := ll.head
    fmt.Print("Linked List: ")
    for current != nil {
        fmt.Printf("%d -> ", current.data)
        current = current.next
    }
    fmt.Println("nil")
}

func (s *Stack) Push(data int) {
    s.items = append(s.items, data)
}

func (s *Stack) Pop() (int, bool) {
    if len(s.items) == 0 {
        fmt.Println("Tidak ada data untuk diproses")
        return 0, false
    }
    index := len(s.items) - 1
    data := s.items[index]
    s.items = s.items[:index]
    return data, true
}

func main() {
    ll := &LinkedList{}
    undoStack := &Stack{}
    redoStack := &Stack{}
    ll.Add(10)
    ll.Add(20)
    ll.Add(30)
    ll.Print()

    ll.Remove(20, undoStack)
    ll.Print()

    ll.Undo(undoStack, redoStack)
    ll.Print()

    ll.Redo(redoStack)
    ll.Print()
}