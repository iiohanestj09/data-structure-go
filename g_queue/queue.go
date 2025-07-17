package main

import "fmt"

type Queue struct {
    items []int
}

func NewQueue() *Queue {
    return &Queue{items: []int{}}
}

func (q *Queue) Enqueue(item int) {
    q.items = append(q.items, item)
    fmt.Printf("Elemen %d ditambahkan ke queue\n", item)
}

func (q *Queue) Dequeue() int {
    if len(q.items) == 0 {
        fmt.Println("Queue kosong!")
        return -1
    }
    front := q.items[0]
    q.items = q.items[1:]
    fmt.Printf("Elemen %d dihapus dari queue\n", front)
    return front
}

func (q *Queue) Display() {
    if len(q.items) == 0 {
        fmt.Println("Queue kosong!")
        return
    }
    fmt.Println("Isi queue (dari depan ke belakang):")
    for i, v := range q.items {
        fmt.Printf("[%d] %d\n", i, v)
    }
}

func main() {
    queue := NewQueue()

    queue.Enqueue(10)
    queue.Enqueue(20)
    queue.Enqueue(30)

    queue.Display()

    queue.Dequeue()

    queue.Display()
}
