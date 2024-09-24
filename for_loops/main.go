package main

import "fmt"

func main() {
	// index := 1

	// for index <= 10 {
	// 	fmt.Println(index)
	// 	index = index + 1
	// }

	// total := 0

	// for nindex := 1; nindex <= 10; nindex++ {

	// 	total += nindex
	// }
	// fmt.Println(total)

	// index := 0

	// var names = [3]string{"ilhan", "enes", "daniş"}

	// for index < len(names) {
	// 	fmt.Println(names[index])
	// 	index++
	// }

	for index := 0; index <= 10; index++ {
		if index == 3 {
			break
		}
		fmt.Println(index)
	}

	for i := 0; i <= 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}

}
