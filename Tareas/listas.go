package main

import "fmt"

type Task struct {
    Title string
    Done  bool
}

type Node struct {
    value Task
    next  *Node
}

type TaskList struct {
    head *Node
}

func (l *TaskList) Add(title string) {
    newNode := &Node{value: Task{Title: title}}

    if l.head == nil {
        l.head = newNode
        return
    }

    temp := l.head
    for temp.next != nil {
        temp = temp.next
    }
    temp.next = newNode
}

func (l *TaskList) Print() {
    temp := l.head
    for temp != nil {
        fmt.Println("-", temp.value.Title, "| Done:", temp.value.Done)
        temp = temp.next
    }
}

func main() {
    list := TaskList{}

    list.Add("Aprender Go")
    list.Add("Hacer tarea")
    list.Add("Estudiar estructuras")

    list.Print()
}