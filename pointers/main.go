package main

import "fmt"

func main() {
	var a int

	var p *int

	a = 10

	p = &a

	fmt.Println(a)
	fmt.Println(p)

	fmt.Println(*p)

	*p = 20

	fmt.Println(a)

	var numberA = 10
	var numberB int

	var numberC *int

	numberB = numberA

	numberC = &numberA

	*numberC = 20

	fmt.Println(numberA, numberB)
}
