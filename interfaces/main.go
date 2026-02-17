package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Saveable interface defines types that can be saved to a file
type Saveable interface {
	Save() string
}

type User struct {
	Name string
}

type Product struct {
	Name  string
	Price float64
}

func (u User) Save() string {
	return fmt.Sprintf("Saving user: %s", u.Name)
}

func (p Product) Save() string {
	return fmt.Sprintf("Saving product: %s", p.Name)
}

// Generic Save function that works with any Saveable type
func SaveData(data Saveable) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return err
	}
	filename := fmt.Sprintf("%T.json", data) // Use the type name as the filename
	return os.WriteFile(filename, jsonData, 0644)
}

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

	user := User{Name: "John"}
	product := Product{Name: "Laptop", Price: 999.99}

	err := SaveData(user)
	if err != nil {
		fmt.Println("Error saving user:", err)
	}

	err = SaveData(product)
	if err != nil {
		fmt.Println("Error saving product:", err)
	}
}
