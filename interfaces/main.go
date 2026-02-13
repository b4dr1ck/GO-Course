package main

import "fmt"

func main() {
	// type assertions
	var a any = 10
	typeInt, ok := a.(int) // check if a is of type int
	if ok {
		fmt.Printf("a is of type int: %v\n", typeInt)
	}

	// type assertions with type-switch
	switch a.(type) {
	case int:
		fmt.Printf("Integer: %v\n", a)
	case float64:
		fmt.Printf("Float: %v\n", a)
	case string:
		fmt.Printf("String: %v\n", a)
	case bool:
		fmt.Printf("bool: %v\n", a)
	default:
		fmt.Print("Unknown Type\n")
	}
}
