package conversion

import (
	"fmt"
	"strconv"
)

func ConvertStringsToFloats(lines []string) ([]float64, error) {

	convLines := []float64{}
	for _, line := range lines {
		floatValue, err := strconv.ParseFloat(line, 64)
		if err != nil {
			fmt.Println("Error converting prices from file")
			return convLines, err
		}
		convLines = append(convLines, floatValue)
	}

	return convLines, nil
}
