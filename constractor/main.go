package main

import "fmt"

type hoge struct {
	name string
	age  int16
}

func main() {
	test := newHoge()
	fmt.Println(test)
	fmt.Println(test.name)
	fmt.Println(test.age)
}

func newHoge() *hoge {
	return &hoge{
		name: "fuganame",
		age:  32,
	}
}
