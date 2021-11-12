package user

import (
	"strings"

	"../model"
)

type Venom struct {
	Name string
	Age  int
}

func NewVenom() Venom {
	return Venom{
		Name: "ヴェノム",
		Age:  4,
	}
}

func (v Venom) Display() string {
	return "My name is Venom"
}

func (v Venom) Wishlist() string {
	return "チョコが食べたい"
}

func (v Venom) CanGet(sack []model.Car) string {
	gds := make([]string, 0)

	for _, gift := range sack {
		if gift.GetColor() == "あか" && gift.GetCategory() == "のりもの" {
			gds = append(gds, gift.Display())
		}
	}

	if len(gds) == 0 {
		return "ほしいものがみつからなかった..."
	}
	return strings.Join(gds, "\n  ")
}
