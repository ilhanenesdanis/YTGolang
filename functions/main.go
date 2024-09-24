package main

import "fmt"

func main() {
	add(10, 20)

	total := sum(19, 22)

	fmt.Println(total)

	calculate, difference := calculation(10, 20)

	fmt.Println(calculate)
	fmt.Println(difference)

	var numbers = []int{10, 20, 2, 3, 5}

	fmt.Println(sumArray(numbers))

	fmt.Println(sumNparameter(1, 2, 3, 5, 634, 12))

}

func sumArray(numbers []int) int {
	total := 0

	for _, value := range numbers {
		total += value
	}
	return total
}

func add(x int, y int) {
	fmt.Println(x + y)
}

func sum(x int, y int) int {
	return x + y

}
func calculation(x int, y int) (int, int) {
	return sum(x, y), x - y
}
func sumNparameter(numbers ...int) int {
	return sumArray(numbers)
}
