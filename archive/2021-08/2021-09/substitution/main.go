package main

import "fmt"

var fuga string = "fuga"
var method = (*hoge).greet

func main() {
	a := NewHoge(fuga)
	a.greet()
	b := NewHoge("test")
	method(b)
}

type hoge struct {
	name string
}

func NewHoge(name string) *hoge {
	return &hoge{name: name}
}
func (h *hoge) greet() {
	fmt.Println("hello!! my name is " + h.name)
}
