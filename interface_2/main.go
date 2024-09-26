package main

import "fmt"

type XEncoder struct {
}

type IEncoder interface {
	Encode(value string)
	Decode(valus string)
}

func (xEncode *XEncoder) Encode(value string) {
	fmt.Println("XEncoder ile encode edildi ")
}
func (xEncode *XEncoder) Decode(value string) {
	fmt.Println("XEncoder ile decode edildi ")
}

func main() {

	var encoder IEncoder

	encoder = &XEncoder{}

	encoder.Encode("123")

	encoder.Decode("123")
}
