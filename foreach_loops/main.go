package main

import "fmt"

func main() {
	var numbers = []int{1, 2, 3, 4}
	//for loops
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}
	//foreach loops

	for index, value := range numbers {
		fmt.Println(index, value)

	}

	var language = "GoLang"

	for _, character := range language {

		fmt.Printf("Character %c\n", character)

	}

	var names = map[string]int{
		"İlhan": 10,
		"Enes":  15,
		"Danis": 20,
	}

	for key, value := range names {

		fmt.Println(key, value)
	}
}
