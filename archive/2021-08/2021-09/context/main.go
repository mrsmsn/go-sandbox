package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "hoge", "value test")

	fmt.Println(ctx.Value("hoge").(string))
}
