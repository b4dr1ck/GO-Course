package main

import (
	"fmt"
	"math/rand/v2"
)

var score int = 0

func main() {
	showHeader()

	for {
		showOptions()

		var userInput int
		fmt.Scan(&userInput)

		switch userInput {
		case 1:
			calcWrapper("+")
		case 2:
			calcWrapper("-")
		case 3:
			calcWrapper("*")
		case 4:
			calcWrapper("/")
		case 5:
			return

		}
	}
}

func calcWrapper(operator string) {
	var userInput int
	var number1 = rand.IntN(100)
	var number2 = rand.IntN(100)
	var result int

	switch operator {
	case "+":
		result = number1 + number2
	case "-":
		result = number1 - number2
	case "*":
		result = number1 * number2
	case "/":
		result = number1 / number2
	}

	fmt.Printf("%v %s %v = ", number1, operator, number2)

	fmt.Scan(&userInput)

	if userInput == result {
		fmt.Println("☑ Correct!")
		score += 1
	} else {
		fmt.Println("☒ Wrong!")
		fmt.Printf("Result is: %v\n", result)
		score -= 1
	}
	fmt.Println()
}

func showHeader() {
	fmt.Println(`
▖  ▖  ▗ ▌   ▄▖    ▘      
▛▖▞▌▀▌▜▘▛▌  ▐ ▛▘▀▌▌▛▌█▌▛▘
▌▝ ▌█▌▐▖▌▌  ▐ ▌ █▌▌▌▌▙▖▌ 
                         `)
}

func showOptions() {
	fmt.Printf("Score: %v\n", score)
	fmt.Println("1) Addition")
	fmt.Println("2) Substraction")
	fmt.Println("3) Multiplication")
	fmt.Println("4) Division")
	fmt.Println("5) Exit")
	fmt.Println("-----------------------")
	fmt.Println("Make your choice (type in number 1-4):")
}
