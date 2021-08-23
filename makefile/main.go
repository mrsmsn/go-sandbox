package main

import "fmt"

var (
	Version  = "x.y.z"
	Revision = "xxx"
)

func main() {
	fmt.Println("Version:" + Version + "  Revision:" + Revision)
}
