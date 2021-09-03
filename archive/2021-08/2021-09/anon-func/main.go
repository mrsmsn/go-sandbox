package main

import "fmt"

func main() {
	func() { fmt.Println("Anon func test2") }()
}
