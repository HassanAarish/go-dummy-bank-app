package fileOps

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func GetFloatFromFile(filaName string, defaultValue float64) (float64, error) {
	data, err := os.ReadFile(filaName)

	if err != nil {
		return defaultValue, errors.New("Failed to find file.")
	}

	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)

	if err != nil {
		return defaultValue, errors.New("Failed to parse file.")
	}

	return value, nil
}

func WriteFloatToFile(value float64, fileName string) {
	valueText := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(valueText), 0664)
}
