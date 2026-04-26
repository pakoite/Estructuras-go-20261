package main

import "fmt"
/* Leer alumnos.md */


// Materia
type Subject struct {
    Name  string
    Grade float64
}

// Alumno
type Student struct {
    Name     string
    Subjects map[string]Subject
}

// Agregar materia
func (s *Student) AddSubject(name string, grade float64) {
    if s.Subjects == nil {
        s.Subjects = make(map[string]Subject)
    }

    s.Subjects[name] = Subject{
        Name:  name,
        Grade: grade,
    }
}

// Promedio
func (s Student) Average() float64 {
    total := 0.0

    for _, sub := range s.Subjects {
        total += sub.Grade
    }

    if len(s.Subjects) == 0 {
        return 0
    }

    return total / float64(len(s.Subjects))
}

// Mostrar info
func (s Student) Print() {
    fmt.Println("Alumno:", s.Name)

    for _, sub := range s.Subjects {
        fmt.Println("-", sub.Name, ":", sub.Grade)
    }

    fmt.Println("Promedio:", s.Average())
}

func main() {
    student := Student{Name: "Carlos"}

    student.AddSubject("Matemáticas", 90)
    student.AddSubject("Programación", 95)
    student.AddSubject("Física", 80)

    student.Print()
}