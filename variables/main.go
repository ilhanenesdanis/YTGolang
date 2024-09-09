package main

import (
	"fmt"
	"reflect"
)

func main() {

	var productName string

	var quantity int

	var discount float32

	var isInStock bool

	productName = "Golang Variables"

	quantity = 10

	discount = 0.37

	isInStock = true

	//reflect.TypeOf(productName) değişken tipini tanımlamak için kullanılır
	fmt.Println(productName, reflect.TypeOf(productName))
	fmt.Println(quantity, reflect.TypeOf(quantity))
	fmt.Println(discount, reflect.TypeOf(discount))
	fmt.Println(isInStock, reflect.TypeOf(isInStock))

	pName := "Golang"

	fmt.Println(pName)

	fmt.Println("ProductName", productName, "Quantity", quantity, "Discount", discount, "isInStock", isInStock)

	fmt.Printf("Product Name %s\n, Quantit: %d\n ,Discount: %f\n", productName, quantity, discount)
}
