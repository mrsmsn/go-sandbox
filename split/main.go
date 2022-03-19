package main

import (
	"fmt"
	"strings"
)

func main() {
	path := "_submodules/gopkg/extract/order/_data/hoeghoge/parser.toml"
	fmt.Println(len(strings.Split(path, "/")))
}
