package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser user.User

	// init a new struct-instance
	appUser, err := user.New(userFirstName, userLastName, userBirthdate)
	if err != nil {
		fmt.Println(err)
		return
	}

	// create an admin
	admin := user.NewAdmin("admin@example.com", "test123")
	admin.User.OutputUserData()

	// call the defined methods
	appUser.OutputUserData()
	fmt.Println("Clearing user name...")
	appUser.ClearUserName()
	appUser.OutputUserData() // show data after clearing
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
