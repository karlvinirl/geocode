package main

import "fmt"

// example of typed function
type transformFnc func(int) int

func main() {

	numbers := []int{1, 2, 3, 4}
	numbers2 := []int{2, 3, 4}
	fmt.Println(numbers)
	fmt.Println(numbers2)

	//example of taking function as value
	tNumbers := transformNumbers(&numbers, getTransformFunction(&numbers))
	fmt.Println(tNumbers)

	tNumbers = transformNumbers(&numbers2, getTransformFunction(&numbers2))
	fmt.Println(tNumbers)

}

// example of returning a function as value
func getTransformFunction(numbers *[]int) transformFnc {
	if (*numbers)[0] == 1 {
		return doubleNumbers
	} else {
		return tripleNumbers
	}
}

// example of taking function as value
func transformNumbers(numbers *[]int, transform transformFnc) []int {
	tNumbers := []int{}

	for _, val := range *numbers {

		tNumbers = append(tNumbers, transform(val))
	}

	return tNumbers
}

func doubleNumbers(number int) int {

	return number * 2
}

func tripleNumbers(number int) int {

	return number * 3
}
