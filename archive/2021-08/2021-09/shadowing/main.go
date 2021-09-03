package main

import "fmt"

var (
	number = 1
)

func main() {
	fmt.Println(&number)
	number := number
	number = 9
	fmt.Println(&number)
}
