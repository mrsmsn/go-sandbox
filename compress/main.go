// package main
//
// import (
// 	"bytes"
// 	"compress/gzip"
// 	"encoding/base64"
// 	"fmt"
// 	"io/ioutil"
// )
//
// func main() {
// 	var b bytes.Buffer
// 	gz := gzip.NewWriter(&b)
// 	if _, err := gz.Write([]byte("YourDataHere")); err != nil {
// 		panic(err)
// 	}
// 	if err := gz.Flush(); err != nil {
// 		panic(err)
// 	}
// 	if err := gz.Close(); err != nil {
// 		panic(err)
// 	}
// 	str := base64.StdEncoding.EncodeToString(b.Bytes())
// 	fmt.Println(str)
// 	data, _ := base64.StdEncoding.DecodeString(str)
// 	fmt.Println(data)
// 	rdata := bytes.NewReader(data)
// 	r, _ := gzip.NewReader(rdata)
// 	s, _ := ioutil.ReadAll(r)
// 	fmt.Println(string(s))
//
// }
//
package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
)

const html = `<!DOCTYPE html>
		<html><head></head><body>
		
		<h1>My First Heading</h1>
		
		<p>My first paragraph.</p>
		
		
		
		</body></html>`

func main() {
	// 圧縮
	var buffer bytes.Buffer
	writer := gzip.NewWriter(&buffer)
	writer.Write([]byte(html))
	writer.Close()
	b := buffer.Bytes()
	fmt.Println(string(b))

	// 展開
	reader, _ := gzip.NewReader(&buffer)
	output := bytes.Buffer{}
	output.ReadFrom(reader)
	s := string(output.Bytes())
	fmt.Println(s)
}
