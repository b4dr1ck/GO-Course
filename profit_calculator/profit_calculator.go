package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {

	fmt.Println("Last Result:")
	lastResult, err := readResultsFromFile()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(lastResult)

	revenue, err := getUserInput("Revenue")
	if err != nil {
		fmt.Println(err)
		panic("Abort Application")
	}

	expenses, err := getUserInput("Expenses")
	if err != nil {
		fmt.Println(err)
		panic("Abort Application")
	}

	taxRate, err := getUserInput("Tax Rate")
	if err != nil {
		fmt.Println(err)
		panic("Abort Application")
	}

	ebt, profit, ratio := calculations(revenue, expenses, taxRate)

	fmt.Println("EBT (Earnings Before Tax):", ebt)
	fmt.Println("Profit:", profit)
	fmt.Printf("EBT to Profit Ratio: %.2f\n", ratio)

	storeResultsInFile(fmt.Sprintf("EBT: %v\nProfit %v\nEBT to Profit Ratio %.2f\n", ebt, profit, ratio))
}

func getUserInput(text string) (float64, error) {
	var userInput float64
	fmt.Print("Enter ", text, ": ")
	fmt.Scan(&userInput)

	if userInput <= 0 {
		return 0, errors.New("Value (" + text + ") must be bigger than zero")
	}

	return userInput, nil
}

func calculations(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	return ebt, profit, ratio
}

func storeResultsInFile(data string) {
	os.WriteFile("results.txt", []byte(data), 0644)
}

func readResultsFromFile() (string, error) {
	results, err := os.ReadFile("results.txt")
	if err != nil {
		return "", errors.New("Can not read file: results.txt")
	}

	return string(results), nil
}
