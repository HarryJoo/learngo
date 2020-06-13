package main

import "fmt"

func superAdd(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}

	return total

	/*
		for number := range numbers {
			fmt.Println(number)
		}

		return 1
	*/

	/*
		for index, number := range numbers {
			fmt.Println(index, number)
		}

		return 1
	*/

	/*
		for i:= 0; i < len(numbers); i++ {
			fmt.Println(i)
		}

		return 1
	*/
}

func main() {
	result := superAdd(1, 2, 3, 4, 5, 6)
	fmt.Println(result)
}
