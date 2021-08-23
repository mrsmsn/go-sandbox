package main

import "fmt"

func main() {
	const sample = "This is a pen"
	for i := 0; i < len(sample); i++ {
		fmt.Printf("%x ", sample[i])
	}
}
