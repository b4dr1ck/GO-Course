package main

import (
	"fmt"

	"example.com/bank/fileops"
	"github.com/Pallinder/go-randomdata"
)

const balanceFile = "balance.txt"

func main() {
	accoutBalance, err := fileops.GetFloatFromFile(balanceFile)
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
	}

	fmt.Println("")
	fmt.Println(` /$$$$$$   /$$$$$$        /$$$$$$$   /$$$$$$  /$$   /$$ /$$   /$$
 /$$__  $$ /$$__  $$      | $$__  $$ /$$__  $$| $$$ | $$| $$  /$$/
| $$  \__/| $$  \ $$      | $$  \ $$| $$  \ $$| $$$$| $$| $$ /$$/ 
| $$ /$$$$| $$  | $$      | $$$$$$$ | $$$$$$$$| $$ $$ $$| $$$$$/  
| $$|_  $$| $$  | $$      | $$__  $$| $$__  $$| $$  $$$$| $$  $$  
| $$  \ $$| $$  | $$      | $$  \ $$| $$  | $$| $$\  $$$| $$\  $$ 
|  $$$$$$/|  $$$$$$/      | $$$$$$$/| $$  | $$| $$ \  $$| $$ \  $$
 \______/  \______/       |_______/ |__/  |__/|__/  \__/|__/  \__/`)
	fmt.Println("Phone: ", randomdata.PhoneNumber())

	for {

		presentOptions()

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
			fileops.WriteFloatToFile(accoutBalance, balanceFile)
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
			fileops.WriteFloatToFile(accoutBalance, balanceFile)
			fmt.Printf("Balance updated! New amount: %.2f€\n", accoutBalance)

		default:
			fmt.Println("Good Bye!")
			fmt.Println("Thanks for choosing our bank")
			return
		}
	}
}
