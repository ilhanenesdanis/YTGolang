package main

import "fmt"

func main() {

	var products []IShippable = []IShippable{
		&Book{Desi: 10},
		&Book{Desi: 20},
		&Electronic{Desi: 20},
		&Flover{Desi: 3},
	}

	fmt.Println(calculateTotalShippingCost(products))

}

func calculateTotalShippingCost(products []IShippable) int {
	total := 0
	for _, product := range products {
		total += product.ShippingCost()
	}
	return total

}

func (electronic *Electronic) ShippingCost() int {
	return 5 + electronic.Desi*3

}
func (book *Book) ShippingCost() int {
	return 5 + book.Desi*3
}
func (flover *Flover) ShippingCost() int {
	return 5 + flover.Desi*3
}

type Book struct {
	Desi int
}
type Electronic struct {
	Desi int
}
type Flover struct {
	Desi int
}
type IShippable interface {
	ShippingCost() int
}
