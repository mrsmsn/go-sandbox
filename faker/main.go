package main

import (
	"fmt"

	"github.com/jaswdr/faker"
)

func main() {
	faker := faker.New()

	fmt.Println(faker.Company().Name())

	fmt.Println(faker.Address().Address())

	fmt.Println(faker.Lorem().Text(100))
}
