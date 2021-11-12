package main

import (
	"errors"
	"fmt"
	"go/format"
	"io/ioutil"
	"os"
	"strings"

	"github.com/cbroglie/mustache"
	"github.com/mrsmsn/go-sandbox/metapro/model"
	"gopkg.in/yaml.v2"
)

type Renderer struct {
	Tmpl    *mustache.Template
	Path    string
	Postfix string
}

var (
	basePath = "./"

	// データ
	inputBasePath = basePath + "yaml/"
	carsInputPath = inputBasePath + "cars.yaml"

	// テンプレ
	templateBasePath = basePath + "templates/"

	carTemplatePath = templateBasePath + "car.mustache"

	outputBasePath = "../app/"

	carOutputPath = outputBasePath + "car/"
)

func main() {

	b, err := ioutil.ReadFile(carsInputPath)
	if !errors.Is(err, nil) {
		panic(err)
	}
	container := model.CarContainer{}
	container.Cars = make([]model.Car, 0)

	err = yaml.Unmarshal(b, &container)
	if !errors.Is(err, nil) {
		panic(err)
	}
	generateCar(container.Cars, carsInputPath)

}

func generateCar(cars []model.Car, path string) {
	template, err := mustache.ParseFile(path)
	if !errors.Is(err, nil) {
		panic(err)
	}
	for _, p := range cars {
		for _, r := range []Renderer{
			{
				Tmpl: template,
				Path: carOutputPath,
			},
		} {
			outputFile(r, p.Name, p)
		}
	}
}

func outputFile(r Renderer, name string, data interface{}) {
	output, err := r.Tmpl.Render(data)
	if !errors.Is(err, nil) {
		panic(err)
	}

	fmt.Println("-----------------")
	fmt.Println(output)
	fmt.Println("-----------------")
	outputBytes, err := format.Source([]byte(output))
	if !errors.Is(err, nil) {
		panic(err)
	}
	// outputBytes := []byte(output)

	// ディレクトリを作成、存在する場合は無視する。
	_ = os.MkdirAll(r.Path, 0755)

	// []byte をファイルに上書きしています。
	filename := r.Path + strings.ToLower(name) + r.Postfix + ".go"
	err = ioutil.WriteFile(filename, outputBytes, 0755)
	if err != nil {
		panic(err)
	}

	fmt.Printf("mustache: generate %s\n", filename)

}
