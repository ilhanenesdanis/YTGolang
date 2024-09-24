package main

import "fmt"

func main() {

	var customer1 = Customer{
		Id:   1,
		Name: "İlhan Enes",
		Age:  23,
		address: Address{
			city:     "İstanbul",
			district: "Pendik",
		},
	}

	var customer2 = Customer{
		Id:   2,
		Name: "Ramazan",
		Age:  52,
		address: Address{
			city:     "İstanbul",
			district: "Maltepe",
		},
	}

	customer1.Age = 25
	fmt.Println(customer1)
	fmt.Println(customer2)

	customer1.ToString()

	customer2.ToString()
	changeName(&customer1)

	customer1.ToString()

	customer1.changeName("İlhan Enes")

	customer1.ToString()

}

type Customer struct {
	Id      int64
	Name    string
	Age     int
	address Address
}
type Address struct {
	city     string
	district string
}

func (customer Customer) ToString() {
	fmt.Printf("Name: %s , Age: %d", customer.Name, customer.Age)
}
func changeName(customer *Customer) {
	customer.Name = "Rümeysa Daniş"
}
func (customer *Customer) changeName(newName string) {
	customer.Name = newName
}
