package main

import "fmt"

type Stack struct {
    items []string
}

func (s *Stack) Push(url string) {
    s.items = append(s.items, url)
}

func (s *Stack) Pop() string {
    if len(s.items) == 0 {
        return ""
    }
    last := s.items[len(s.items)-1]
    s.items = s.items[:len(s.items)-1]
    return last
}

// 🔹 Modelo real
type Browser struct {
    current string
    history Stack
}

func (b *Browser) Visit(url string) {
    if b.current != "" {
        b.history.Push(b.current)
    }
    b.current = url
}

func (b *Browser) Back() {
    prev := b.history.Pop()
    if prev != "" {
        b.current = prev
    }
}

func main() {
    b := Browser{}

    b.Visit("google.com")
    b.Visit("github.com")
    b.Visit("stackoverflow.com")

    fmt.Println("Actual:", b.current)

    b.Back()
    fmt.Println("Después de back:", b.current)
}