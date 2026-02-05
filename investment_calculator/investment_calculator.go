package main

import (
	"fmt"
	"math"
)

const inflationRate float64 = 2.5

func main() {
	var investmentAmount float64
	var years float64
	var expectedReturnRate float64 = 5.5

	fmt.Print("Enter investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter number of years: ")
	fmt.Scan(&years)

	fmt.Print("Enter expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, years, expectedReturnRate)

	formattedFV := fmt.Sprintf("%.2f", futureValue)
	formattedFRV := fmt.Sprintf("%.2f", futureRealValue)

	fmt.Println("=============================")
	fmt.Println("Future Value:", formattedFV)
	fmt.Printf("Future Real Value (adjusted for %.2f%% inflation): %s\n", inflationRate, formattedFRV)
	fmt.Println("=============================")
}

func calculateFutureValues(investmentAmount, years, expectedReturnRate float64) (float64, float64) {
	fv := investmentAmount * math.Pow((1+expectedReturnRate/100), years)
	frv := fv / math.Pow((1+inflationRate/100), years)

	return fv, frv
}

// Alternative implementation with named return values
func calculateFutureValues2(investmentAmount, years, expectedReturnRate float64) (fv float64, frv float64) {
	fv = investmentAmount * math.Pow((1+expectedReturnRate/100), years)
	frv = fv / math.Pow((1+inflationRate/100), years)
	return
}
