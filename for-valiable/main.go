package main

import "fmt"

type hoge struct {
	Name string
	Age  int32
}

func main() {
	list := []hoge{
		{
			Name: "test",
			Age:  15,
		},
		{
			Name: "test2",
			Age:  18,
		},
		{
			Name: "test3",
			Age:  21,
		},
	}

	for _, v := range list {
		var fuga []*hoge
		if v.Name == "test2" {
			fuga = append(fuga, &hoge{
				Name: "初期化",
				Age:  99,
			})
		}
		// 毎回初期化されていることを確認
		fmt.Printf("%s\n", fuga)
	}

	/* result output */
	// []
	// [%!s(*main.hoge=&{初期化 99})]
	// []
	/* result output end */

}
