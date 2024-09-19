package main

import "fmt"

func main() {
	// var names [3]string

	// names[0] = "ilhan"
	// names[1] = "rümeysa"
	// names[2] = "ramazan"

	// fmt.Println(names)

	var names = [4]string{"ilhan", "rümeysa", "ramazan", "ecrin"}
	fmt.Println(names)

	names[1] = "enes"

	fmt.Println(names[0:2])

	var nameList = []string{"ilhan", "rümeysa", "ramazan", "enes"}

	fmt.Println(nameList)
	nameList = append(nameList, "Serhat")
	fmt.Println(nameList)

}
