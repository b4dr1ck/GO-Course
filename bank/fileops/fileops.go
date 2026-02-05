package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func GetFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)

	if err != nil {
		return 0.0, errors.New("file not found")
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		return 0.0, errors.New("value can't be converted to float")
	}

	return balance, nil
}

func WriteFloatToFile(value float64, fileName string) {
	newValue := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(newValue), 0644)
}
