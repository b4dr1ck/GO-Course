package main

import "fmt"

func main() {
	var productNames [4]string
	productNames = [4]string{"Book"}
	productNames[2] = "Carpet"
	prices := [4]float64{10.99, 39.99, 50.0, 12.99}

	// can also use with mixed data-types
	mixedData := [5]any{1, 5.0, "hello"}

	// dynamic array
	amounts := []int{1, 2, 5, 6, 3, 1}
	amounts = append(amounts, 12)

	// remove first element (with slices)
	amounts = amounts[1:]

	// slices (start-index is included, end-index is excluded)
	featuredPrices := prices[1:3]

	fmt.Printf("prices: %v\n", prices)
	fmt.Printf("featuredPrices (prices[1:3]): %v\n", featuredPrices)
	fmt.Printf("productNames: %v\n", productNames)
	fmt.Printf("The price of the %s is %.2f€ \n", productNames[2], prices[2])
	fmt.Printf("mixedData: %v\n", mixedData)
	fmt.Printf("amounts: %v\n", amounts)

}
