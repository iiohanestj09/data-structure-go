package main

import "fmt"

type Node struct {
    data int
    next *Node
}

type LinkedList struct {
    head *Node
}

func (list *LinkedList) Insert(value int) {
    newNode := &Node{data: value}
    if list.head == nil {
        list.head = newNode
    } else {
        temp := list.head
        for temp.next != nil {
            temp = temp.next
        }
        temp.next = newNode
    }
}

func (list *LinkedList) Delete(value int) {
    if list.head == nil {
        fmt.Println("Linked List kosong")
        return
    }

    if list.head.data == value {
        list.head = list.head.next
        return
    }

    temp := list.head
    for temp.next != nil && temp.next.data != value {
        temp = temp.next
    }

    if temp.next == nil {
        fmt.Println("Data tidak ditemukan dalam Linked List")
    } else {
        temp.next = temp.next.next
    }
}

func (list *LinkedList) Display() {
    if list.head == nil {
        fmt.Println("Linked List kosong")
        return
    }

    temp := list.head
    for temp != nil {
        fmt.Print(temp.data, " -> ")
        temp = temp.next
    }
    fmt.Println("nil")
}

func main() {
    list := LinkedList{}

    list.Insert(10)
    list.Insert(20)
    list.Insert(30)

    fmt.Println("Isi Linked List sekarang:")
    list.Display()

    list.Delete(20)
    fmt.Println("Isi Linked List setelah delete:")
    list.Display()
}
