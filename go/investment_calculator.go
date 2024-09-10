package main

import (
	"fmt"
	"math"
)

func main() {

	const inflationRate = 2.5
	var investmentAmount float64
	expectedReturnRate := 5.5
	years := 10.0

	fmt.Print("Investment amount:")
	fmt.Scan(&investmentAmount)
	fmt.Print("Expected Return Rate:")
	fmt.Scan(&expectedReturnRate)
	fmt.Print("Years of investment:")
	fmt.Scan(&years)

	//example of default assignment and type
	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)

	//example of Sprintf usage.
	futureValueFmt := fmt.Sprintf("Future Value: %v\n", futureValue)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Print(futureValueFmt)

	//example of formatting using printf
	fmt.Printf("Future real value: %0.3f\n", futureRealValue)

}
