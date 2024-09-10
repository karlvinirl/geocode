package fileops

import (
	"fmt"
	"os"
	"strconv"
)

func WriteFloatValue(filename string, value float64) {
	valueString := fmt.Sprint(value)
	os.WriteFile(filename, []byte(valueString), 0644)
}

func ReadFloatValue(filename string) float64 {
	valueByte, _ := os.ReadFile(filename)
	valueText := string(valueByte)
	value, _ := strconv.ParseFloat(valueText, 64)
	return value
}
