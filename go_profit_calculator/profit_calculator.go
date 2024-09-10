package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {

	var revenue, expenses, taxRate float64
	var revenueError, expensesError, taxRateError error
	fmt.Println("Welcome to profit Calculator")

	revenue, revenueError = getUserInput("Enter revenue:")
	if revenueError != nil {
		fmt.Println(revenueError)
		return
	}
	expenses, expensesError = getUserInput("Enter expenses:")
	if expensesError != nil {
		fmt.Println(expensesError)
		return
	}
	taxRate, taxRateError = getUserInput("Tax Rate:")
	if taxRateError != nil {
		fmt.Println(taxRateError)
		return
	}

	eat, ebt, ratio := calculateValues(revenue, expenses, taxRate)
	displayAnswer(eat, ebt, ratio)
	writeToFile(eat, ebt, ratio)

}

func calculateValues(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	eat := ebt - (ebt / 100 * taxRate)
	ratio := ebt / eat

	return eat, ebt, ratio
}

func getUserInput(prompt string) (float64, error) {
	fmt.Println(prompt)
	var userInput float64
	fmt.Scan(&userInput)
	if userInput <= 0 {
		return 0, errors.New("Input number must be greater than 0")
	}

	return userInput, nil
}

func displayAnswer(ebt, eat, ratio float64) {
	fmt.Println("Earnings before tax: ", ebt)
	fmt.Println("Earnings after tax:", eat)
	fmt.Println("Ratio of EBT/Profit: ", ratio)
}

func writeToFile(ebt, eat, ratio float64) {
	resultsText := fmt.Sprintf("EBT: %.1f\nEAT: %.1f\nRatio: %1.f\n", ebt, eat, ratio)
	os.WriteFile(("ProfitResults.txt"), []byte(resultsText), 0644)
}
