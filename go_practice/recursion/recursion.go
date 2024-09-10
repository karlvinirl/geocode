package main

import "fmt"

func main() {

	fmt.Println(calcFactorial(5))
}

func calcFactorial(number int) int {
	// recursive example
	if number == 0 {
		return 1
	}

	return number * calcFactorial(number-1)

	//  standard solution.
	//	total := 1
	//	for i := 1 ; i <= number ; i++ {
	//		total = total*i
	//		fmt.Println(total)
	//
	//	}
	//	return total

}
