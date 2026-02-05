package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const balanceFile = "balance.txt"
const defaultBalance = 1000.00

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(balanceFile)

	if err != nil {
		return defaultBalance, errors.New("balance file not found")
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		return defaultBalance, errors.New("value can't be converted to float")
	}

	return balance, nil
}

func writeBalanceToFile(balance float64) {
	newBalance := fmt.Sprint(balance)
	os.WriteFile(balanceFile, []byte(newBalance), 0644)
}

func main() {
	accoutBalance, err := getBalanceFromFile()
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("------------------------")
		panic("Abort Application")
	}

	fmt.Println("Welcome to 'Go Bank'!")
	fmt.Println("----------------------------")

	for {

		fmt.Println("What do you want to do?")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Printf("Your balance is %.2f€\n", accoutBalance)
		case 2:
			fmt.Print("Your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Deposit amount must be positive!")
				continue
			}

			accoutBalance += depositAmount
			writeBalanceToFile(accoutBalance)
			fmt.Printf("Balance updated! New amount: %.2f€\n", accoutBalance)
		case 3:
			fmt.Print("Your withdrawal: ")
			var whithdrawalAmount float64
			fmt.Scan(&whithdrawalAmount)

			if whithdrawalAmount <= 0 {
				fmt.Println("Withdrawal amount must be positive!")
				continue
			}

			if whithdrawalAmount > accoutBalance {
				fmt.Println("Insufficient funds!")
				continue
			}

			accoutBalance -= whithdrawalAmount
			writeBalanceToFile(accoutBalance)
			fmt.Printf("Balance updated! New amount: %.2f€\n", accoutBalance)

		default:
			fmt.Println("Good Bye!")
			fmt.Println("Thanks for choosing our bank")
			return
		}
	}
}
