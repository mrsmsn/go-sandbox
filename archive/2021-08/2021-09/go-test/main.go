package main

import (
	"fmt"

	"github.com/mrsmsn/go-sandbox/go-test/hello"
)

func main() {
	s := hello.GetHello("山澤さん")
	fmt.Println(s)
}
