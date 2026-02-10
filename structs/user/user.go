package user

import (
	"errors"
	"fmt"
	"regexp"
	"time"
)

// initialize a struct
type User struct {
	firstName    string
	lastName     string
	birthdate    string
	creationTime time.Time
}

// add a method to a struct (receiver)
func (u User) OutputUserData() {
	fmt.Println(u.firstName, u.lastName, u.birthdate)
}

// must use a pointer in the receiver to mutate the original struct
func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}

// constructor for the struct (not a feature just a design-pattern)
func New(firstName, lastName, birthdate string) (User, error) {
	// validation
	if firstName == "" || lastName == "" || birthdate == "" {
		return User{}, errors.New("Input may not be empty!")
	}

	match, _ := regexp.MatchString(`^(0[1-9]|1[0-2])/(0[1-9]|[12][0-9]|3[01])/\d{4}$`, birthdate)
	if !match {
		return User{}, errors.New("Birthdate must be in the format MM/DD/YYYY!")
	}

	// set the values to the struct and return it
	return User{
		firstName:    firstName,
		lastName:     lastName,
		birthdate:    birthdate,
		creationTime: time.Now(),
	}, nil
}
