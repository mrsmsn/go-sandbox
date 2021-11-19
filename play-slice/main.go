package main

import "fmt"

func main() {
	list := []int{}

	// 当たり前だが空のスライスでfor文回しても実行されない
	for _, v := range list {
		fmt.Println(v)
	}

	fmt.Println("aaaa")
}
