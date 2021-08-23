package main

import "fmt"

func main() {
	apple := "りんご"
	fmt.Println(apple + "アドレス")
	fmt.Println(&apple)
	print(apple, nil)
}

func print(word1 string, word4 *string) {
	fmt.Println("------関数スタート-----")
	fmt.Println(word1 + "アドレス")
	fmt.Println(&word1)
	if word4 != nil {
		fmt.Println(*word4 + "アドレス")
		fmt.Println(word4)
	}
	fmt.Println("------関数エンド-----")
}
