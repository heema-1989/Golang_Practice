package main

import "fmt"

type Book struct {
	title, author string
}
type Magazine struct {
	title string
	issue int
}

func (b *Book) Assign(t, a string) {
	b.author = a
	b.title = t
}
func (b *Book) Print() {
	fmt.Printf("The title is %s and author is %s\n", b.title, b.author)
}
func (m *Magazine) Assign(t string, i int) {
	m.title = t
	m.issue = i
}
func (m *Magazine) Print() {
	fmt.Printf("The title is %s and issue is %d\n", m.title, m.issue)
}

type Print interface {
	Print()
}

func main() {
	var b Book
	var m Magazine
	b.Assign("Circe", "Madeline Miller")
	m.Assign("Rabbit", 36)
	var i Print
	fmt.Println("Calling interface")
	i = &b
	i.Print()
	i = &m
	i.Print()
}
