package main

import "fmt"

func main() {
	var age int = 32
	var agePointer *int = &age

	fmt.Println(*agePointer)

	age = 22

	fmt.Println(*agePointer)

}
