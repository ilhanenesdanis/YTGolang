package main

import "fmt"

func main() {

	book := &Book{Desi: 25}

	cost := book.ShippingCost()

	fmt.Println(cost)

	books := []Book{
		{Desi: 50},
		{Desi: 20},
		{Desi: 10},
	}
	fmt.Println(calculateTotalShippingCost(books))

	electronic1 := &Electronic{Desi: 20}

	fmt.Println(electronic1.ShippingCost())

	electronics := []Electronic{{Desi: 10}, {Desi: 20}}
	calculateTotalShippingCostOfElectronics(electronics)

	fmt.Println("İnterfaces")

	var product IShippable

	product = &Book{Desi: 10}

	fmt.Println(product.ShippingCost())

	product = &Electronic{Desi: 12}

	fmt.Println(product.ShippingCost())

}

func calculateTotalShippingCost(books []Book) int {
	total := 0
	for _, book := range books {
		total += book.ShippingCost()
	}
	return total

}
func calculateTotalShippingCostOfElectronics(electronics []Electronic) int {
	total := 0
	for _, electronic := range electronics {
		total += electronic.ShippingCost()
	}
	return total

}

func (electronic *Electronic) ShippingCost() int {
	return 10 + electronic.Desi*3

}
func (book *Book) ShippingCost() int {
	return 5 + book.Desi*2
}

type Book struct {
	Desi int
}
type Electronic struct {
	Desi int
}
type IShippable interface {
	ShippingCost() int
}
