package main

import "fmt"

func main() {
	var names map[string]int

	names = make(map[string]int, 0)
	names["İlhan"] = 1
	names["Enes"] = 2
	names["Danis"] = 3

	fmt.Println(names)

	fmt.Println(names["check"])
	fmt.Println(names["Enes"])

	nameList := map[string]int{
		"İlhan": 1,
		"Enes":  2,
		"Danis": 3,
	}

	delete(nameList, "Enes")

	fmt.Println(nameList)

}
