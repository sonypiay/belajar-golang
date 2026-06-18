package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}

func main() {
	result := sumAll(10, 90, 80, 70, 60)
	fmt.Println(result)

	numbers := []int{10, 20, 30, 40, 50}
	result = sumAll(numbers...)

	fmt.Println(result)
}
