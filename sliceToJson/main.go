package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Car struct {
	Name  string `json:"name"`
	Speed int32  `json:"speed"`
}

func main() {
	http.HandleFunc("/", handler) // ハンドラを登録してウェブページを表示させる
	http.ListenAndServe(":3000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	axela := newCar("アクセラ", 60)
	cx5 := newCar("CX5", 80)
	var cars []Car
	cars = append(cars, *axela)
	cars = append(cars, *cx5)
	j, _ := json.Marshal(cars)
	w.Write(j)
	fmt.Println([]byte(j))
}

func newCar(name string, speed int32) *Car {
	return &Car{
		Name:  name,
		Speed: speed,
	}
}
