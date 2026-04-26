package main

import "fmt"

type Client struct {
    Name string
}

type Queue struct {
    items []Client
}

func (q *Queue) Enqueue(c Client) {
    q.items = append(q.items, c)
}

func (q *Queue) Dequeue() Client {
    if len(q.items) == 0 {
        return Client{}
    }
    first := q.items[0]
    q.items = q.items[1:]
    return first
}

// 🔹 Sistema real
type Bank struct {
    line Queue
}

func (b *Bank) Arrive(name string) {
    b.line.Enqueue(Client{Name: name})
}

func (b *Bank) Attend() {
    client := b.line.Dequeue()
    if client.Name != "" {
        fmt.Println("Atendiendo a:", client.Name)
    } else {
        fmt.Println("No hay clientes")
    }
}

func main() {
    bank := Bank{}

    bank.Arrive("Ana")
    bank.Arrive("Luis")
    bank.Arrive("Pedro")

    bank.Attend()
    bank.Attend()
}