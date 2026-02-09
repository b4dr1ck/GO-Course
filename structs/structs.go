package main

import (
	"errors"
	"fmt"
	"time"
)

// initialize a struct
type user struct {
	firstName    string
	lastName     string
	birthdate    string
	creationTime time.Time
}

// add a method to a struct (receiver)
func (u user) outputUserData() {
	fmt.Println(u.firstName, u.lastName, u.birthdate)
}

// must use a pointer in the receiver to mutate the original struct
func (u *user) clearUserName() {
	u.firstName = ""
	u.lastName = ""
}

// constructor for the struct (not a feature just a design-pattern)
func newUser(firstName, lastName, birthdate string) (user, error) {
	// validation
	if firstName == "" || lastName == "" || birthdate == "" {
		return user{}, errors.New("Input may not be empty!")
	}

	// set the values to the struct and return it
	return user{
		firstName:    firstName,
		lastName:     lastName,
		birthdate:    birthdate,
		creationTime: time.Now(),
	}, nil
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser user

	// init a new struct-instance
	appUser, err := newUser(userFirstName, userLastName, userBirthdate)
	if err != nil {
		fmt.Println(err)
		return
	}

	// call the defined methods
	appUser.outputUserData()
	appUser.clearUserName()
	appUser.outputUserData() // show data after clearing
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
